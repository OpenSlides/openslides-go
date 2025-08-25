
-- schema_relational.sql for initial database setup OpenSlides
-- Code generated. DO NOT EDIT.
-- MODELS_YML_CHECKSUM = '01253fa5beea1f850a6806f9b85e26e8'


-- Database parameters

-- Do not log messages lower than WARNING
-- For client side logging this can be overwritten using
--
-- SET client_min_messages TO NOTICE;
--
-- to get the log messages in the client locally.
SET log_min_messages TO WARNING;


-- Function and meta table definitions

CREATE EXTENSION hstore;  -- included in standard postgres-installations, check for alpine

CREATE FUNCTION check_not_null_for_relation_lists() RETURNS trigger as $not_null_trigger$
-- usage with 3 parameters IN TRIGGER DEFINITION:
-- table_name of field to check, usually a field in a view
-- column_name of field to check
-- foreign_key field name of triggered table, that will be used to SELECT the values to check the not null.
DECLARE
    table_name TEXT;
    column_name TEXT;
    foreign_key TEXT;
    foreign_id INTEGER;
    counted INTEGER;
begin
    table_name = TG_ARGV[0];
    column_name = TG_ARGV[1];
    foreign_key = TG_ARGV[2];

    IF (TG_OP = 'INSERT') THEN
        foreign_id := hstore(NEW) -> foreign_key;
        IF (foreign_id is NOT NULL) THEN
            foreign_id = NULL; -- no need to ask DB
        END IF;
    ELSIF (TG_OP = 'UPDATE') THEN
        foreign_id := hstore(NEW) -> foreign_key;
        IF (foreign_id is NULL) THEN
            foreign_id = OLD.used_as_default_projector_for_topic_in_meeting_id;
        END IF;
    ELSIF (TG_OP = 'DELETE') THEN
        foreign_id := hstore(OLD) -> foreign_key;
    END IF;

    IF (foreign_id IS NOT NULL) THEN
        EXECUTE format('SELECT array_length(%I, 1) FROM %I where id = %s', column_name, table_name, foreign_id) INTO counted;
        IF (counted is NULL) THEN
            RAISE EXCEPTION 'Trigger % Exception: NOT NULL CONSTRAINT VIOLATED for %.%', TG_NAME, table_name, column_name;
        END IF;
    END IF;
    RETURN NULL;  -- AFTER TRIGGER needs no return
end;
$not_null_trigger$ language plpgsql;

CREATE FUNCTION log_modified_models() RETURNS trigger AS $log_modified_trigger$
DECLARE
    escaped_table_name varchar;
    operation_var TEXT;
    fqid_var TEXT;
BEGIN
    escaped_table_name := TG_ARGV[0];
    operation_var := LOWER(TG_OP);
    fqid_var :=  escaped_table_name || '/' || NEW.id;
    IF (TG_OP = 'DELETE') THEN
        fqid_var = escaped_table_name || '/' || OLD.id;
    END IF;

    INSERT INTO os_notify_log_t (operation, fqid, xact_id, timestamp)
    VALUES (operation_var, fqid_var, pg_current_xact_id(), 'now')
    ON CONFLICT (operation,fqid,xact_id) DO NOTHING;
    RETURN NULL;  -- AFTER TRIGGER needs no return
END;
$log_modified_trigger$ LANGUAGE plpgsql;

CREATE FUNCTION check_unique_ids_pair()
RETURNS trigger
AS $unique_ids_pair_trigger$
-- usage with 1 parameter IN TRIGGER DEFINITION:
-- base_column_name: name of write fields before adding numeric suffixes
-- Guards against mirrored duplicates by skipping one of the pairs.
DECLARE
    base_column_name text;
    value_1 integer;
    value_2 integer;
BEGIN
    base_column_name = TG_ARGV[0];
    value_1 := hstore(NEW) -> (base_column_name || '_1');
    value_2 := hstore(NEW) -> (base_column_name || '_2');

    IF (value_1 > value_2) THEN
        RETURN NULL;
    END IF;

    RETURN NEW;
END;
$unique_ids_pair_trigger$
LANGUAGE plpgsql;

CREATE FUNCTION notify_transaction_end() RETURNS trigger AS $notify_trigger$
DECLARE
    payload TEXT;
    body_content_text TEXT;
BEGIN
    -- Running the trigger for the first time in a transaction creates the table and after commiting the transaction the table is dropped.
    -- Every next run of the trigger in this transaction raises a notice that the table exists. Setting the log_min_messages to notice increases the noise because of such messages.
    CREATE LOCAL TEMPORARY TABLE
    IF NOT EXISTS tbl_notify_counter_tx_once (
        "id" integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
    ) ON COMMIT DROP;

    -- If running for the first time, the transaction id is send via os_notify.
    IF NOT EXISTS (SELECT * FROM tbl_notify_counter_tx_once) THEN
        INSERT INTO tbl_notify_counter_tx_once DEFAULT VALUES;
        payload := '{"xactId":' ||
            pg_current_xact_id() ||
            '}';
        PERFORM pg_notify('os_notify', payload);
    END IF;

    RETURN NULL;  -- AFTER TRIGGER needs no return
END;
$notify_trigger$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION log_modified_related_models()
RETURNS trigger AS $log_modified_related_trigger$
DECLARE
    operation_var TEXT;
    fqid_var TEXT;
    ref_column TEXT;
    foreign_table TEXT;
    foreign_id TEXT;
    i INTEGER := 0;
BEGIN
    operation_var := LOWER(TG_OP);

    WHILE i < TG_NARGS LOOP
        foreign_table := TG_ARGV[i];
        ref_column := TG_ARGV[i+1];

        IF (TG_OP = 'DELETE') THEN
            EXECUTE format('SELECT ($1).%I', ref_column) INTO foreign_id USING OLD;
        ELSE
            EXECUTE format('SELECT ($1).%I', ref_column) INTO foreign_id USING NEW;
        END IF;

        IF foreign_id IS NOT NULL THEN
            fqid_var := foreign_table || '/' || foreign_id;
            INSERT INTO os_notify_log_t  (operation, fqid, xact_id, timestamp)
            VALUES (operation_var, fqid_var, pg_current_xact_id(), now())
            ON CONFLICT (operation,fqid,xact_id) DO NOTHING;
        END IF;

        i := i + 2;
    END LOOP;

    RETURN NULL;  -- AFTER TRIGGER needs no return
END;
$log_modified_related_trigger$ LANGUAGE plpgsql;

CREATE TABLE os_notify_log_t (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    operation varchar(32),
    fqid varchar(256) NOT NULL,
    xact_id xid8,
    timestamp timestamptz,
    CONSTRAINT unique_fqid_xact_id_operation UNIQUE (operation,fqid,xact_id)
);


-- Type definitions


-- Table definitions

CREATE TABLE organization_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256),
    description text,
    legal_notice text,
    privacy_policy text,
    login_text text,
    reset_password_verbose_errors boolean,
    enable_electronic_voting boolean,
    enable_chat boolean,
    limit_of_meetings integer CONSTRAINT minimum_limit_of_meetings CHECK (limit_of_meetings >= 0) DEFAULT 0,
    limit_of_users integer CONSTRAINT minimum_limit_of_users CHECK (limit_of_users >= 0) DEFAULT 0,
    default_language varchar(256) CONSTRAINT enum_organization_default_language CHECK (default_language IN ('en', 'de', 'it', 'es', 'ru', 'cs', 'fr')) DEFAULT 'en',
    require_duplicate_from boolean,
    enable_anonymous boolean,
    saml_enabled boolean,
    saml_login_button_text varchar(256) DEFAULT 'SAML login',
    saml_attr_mapping jsonb,
    saml_metadata_idp text,
    saml_metadata_sp text,
    saml_private_key text,
    theme_id integer NOT NULL UNIQUE,
    users_email_sender varchar(256) DEFAULT 'OpenSlides',
    users_email_replyto varchar(256),
    users_email_subject varchar(256) DEFAULT 'OpenSlides access data',
    users_email_body text DEFAULT 'Dear {name},

this is your personal OpenSlides login:

{url}
Username: {username}
Password: {password}


This email was generated automatically.',
    url varchar(256) DEFAULT 'https://example.com'
);



comment on column organization_t.limit_of_meetings is 'Maximum of active meetings for the whole organization. 0 means no limitation at all';
comment on column organization_t.limit_of_users is 'Maximum of active users for the whole organization. 0 means no limitation at all';


CREATE TABLE user_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    username varchar(256) NOT NULL,
    member_number varchar(256),
    saml_id varchar(256) CONSTRAINT minlength_saml_id CHECK (char_length(saml_id) >= 1),
    pronoun varchar(32),
    title varchar(256),
    first_name varchar(256),
    last_name varchar(256),
    is_active boolean DEFAULT True,
    is_physical_person boolean DEFAULT True,
    password varchar(256),
    default_password varchar(256),
    can_change_own_password boolean DEFAULT True,
    email varchar(256),
    default_vote_weight decimal(16,6) CONSTRAINT minimum_default_vote_weight CHECK (default_vote_weight >= 0.000001) DEFAULT '1.000000',
    last_email_sent timestamptz,
    is_demo_user boolean,
    last_login timestamptz,
    external boolean,
    gender_id integer,
    organization_management_level varchar(256) CONSTRAINT enum_user_organization_management_level CHECK (organization_management_level IN ('superadmin', 'can_manage_organization', 'can_manage_users')),
    home_committee_id integer,
    organization_id integer GENERATED ALWAYS AS (1) STORED NOT NULL
);



comment on column user_t.saml_id is 'unique-key from IdP for SAML login';
comment on column user_t.organization_management_level is 'Hierarchical permission level for the whole organization.';


CREATE TABLE meeting_user_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    comment text,
    number varchar(256),
    about_me text,
    vote_weight decimal(16,6) CONSTRAINT minimum_vote_weight CHECK (vote_weight >= 0.000001),
    locked_out boolean,
    user_id integer NOT NULL,
    meeting_id integer NOT NULL,
    vote_delegated_to_id integer
);




CREATE TABLE gender_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    organization_id integer GENERATED ALWAYS AS (1) STORED NOT NULL
);



comment on column gender_t.name is 'unique';


CREATE TABLE organization_tag_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    color varchar(7) CHECK (color is null or color ~* '^#[a-f0-9]{6}$') NOT NULL,
    organization_id integer GENERATED ALWAYS AS (1) STORED NOT NULL
);




CREATE TABLE theme_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    accent_100 varchar(7) CHECK (accent_100 is null or accent_100 ~* '^#[a-f0-9]{6}$'),
    accent_200 varchar(7) CHECK (accent_200 is null or accent_200 ~* '^#[a-f0-9]{6}$'),
    accent_300 varchar(7) CHECK (accent_300 is null or accent_300 ~* '^#[a-f0-9]{6}$'),
    accent_400 varchar(7) CHECK (accent_400 is null or accent_400 ~* '^#[a-f0-9]{6}$'),
    accent_50 varchar(7) CHECK (accent_50 is null or accent_50 ~* '^#[a-f0-9]{6}$'),
    accent_500 varchar(7) CHECK (accent_500 is null or accent_500 ~* '^#[a-f0-9]{6}$') DEFAULT '#2196f3',
    accent_600 varchar(7) CHECK (accent_600 is null or accent_600 ~* '^#[a-f0-9]{6}$'),
    accent_700 varchar(7) CHECK (accent_700 is null or accent_700 ~* '^#[a-f0-9]{6}$'),
    accent_800 varchar(7) CHECK (accent_800 is null or accent_800 ~* '^#[a-f0-9]{6}$'),
    accent_900 varchar(7) CHECK (accent_900 is null or accent_900 ~* '^#[a-f0-9]{6}$'),
    accent_a100 varchar(7) CHECK (accent_a100 is null or accent_a100 ~* '^#[a-f0-9]{6}$'),
    accent_a200 varchar(7) CHECK (accent_a200 is null or accent_a200 ~* '^#[a-f0-9]{6}$'),
    accent_a400 varchar(7) CHECK (accent_a400 is null or accent_a400 ~* '^#[a-f0-9]{6}$'),
    accent_a700 varchar(7) CHECK (accent_a700 is null or accent_a700 ~* '^#[a-f0-9]{6}$'),
    primary_100 varchar(7) CHECK (primary_100 is null or primary_100 ~* '^#[a-f0-9]{6}$'),
    primary_200 varchar(7) CHECK (primary_200 is null or primary_200 ~* '^#[a-f0-9]{6}$'),
    primary_300 varchar(7) CHECK (primary_300 is null or primary_300 ~* '^#[a-f0-9]{6}$'),
    primary_400 varchar(7) CHECK (primary_400 is null or primary_400 ~* '^#[a-f0-9]{6}$'),
    primary_50 varchar(7) CHECK (primary_50 is null or primary_50 ~* '^#[a-f0-9]{6}$'),
    primary_500 varchar(7) CHECK (primary_500 is null or primary_500 ~* '^#[a-f0-9]{6}$') DEFAULT '#317796',
    primary_600 varchar(7) CHECK (primary_600 is null or primary_600 ~* '^#[a-f0-9]{6}$'),
    primary_700 varchar(7) CHECK (primary_700 is null or primary_700 ~* '^#[a-f0-9]{6}$'),
    primary_800 varchar(7) CHECK (primary_800 is null or primary_800 ~* '^#[a-f0-9]{6}$'),
    primary_900 varchar(7) CHECK (primary_900 is null or primary_900 ~* '^#[a-f0-9]{6}$'),
    primary_a100 varchar(7) CHECK (primary_a100 is null or primary_a100 ~* '^#[a-f0-9]{6}$'),
    primary_a200 varchar(7) CHECK (primary_a200 is null or primary_a200 ~* '^#[a-f0-9]{6}$'),
    primary_a400 varchar(7) CHECK (primary_a400 is null or primary_a400 ~* '^#[a-f0-9]{6}$'),
    primary_a700 varchar(7) CHECK (primary_a700 is null or primary_a700 ~* '^#[a-f0-9]{6}$'),
    warn_100 varchar(7) CHECK (warn_100 is null or warn_100 ~* '^#[a-f0-9]{6}$'),
    warn_200 varchar(7) CHECK (warn_200 is null or warn_200 ~* '^#[a-f0-9]{6}$'),
    warn_300 varchar(7) CHECK (warn_300 is null or warn_300 ~* '^#[a-f0-9]{6}$'),
    warn_400 varchar(7) CHECK (warn_400 is null or warn_400 ~* '^#[a-f0-9]{6}$'),
    warn_50 varchar(7) CHECK (warn_50 is null or warn_50 ~* '^#[a-f0-9]{6}$'),
    warn_500 varchar(7) CHECK (warn_500 is null or warn_500 ~* '^#[a-f0-9]{6}$') DEFAULT '#f06400',
    warn_600 varchar(7) CHECK (warn_600 is null or warn_600 ~* '^#[a-f0-9]{6}$'),
    warn_700 varchar(7) CHECK (warn_700 is null or warn_700 ~* '^#[a-f0-9]{6}$'),
    warn_800 varchar(7) CHECK (warn_800 is null or warn_800 ~* '^#[a-f0-9]{6}$'),
    warn_900 varchar(7) CHECK (warn_900 is null or warn_900 ~* '^#[a-f0-9]{6}$'),
    warn_a100 varchar(7) CHECK (warn_a100 is null or warn_a100 ~* '^#[a-f0-9]{6}$'),
    warn_a200 varchar(7) CHECK (warn_a200 is null or warn_a200 ~* '^#[a-f0-9]{6}$'),
    warn_a400 varchar(7) CHECK (warn_a400 is null or warn_a400 ~* '^#[a-f0-9]{6}$'),
    warn_a700 varchar(7) CHECK (warn_a700 is null or warn_a700 ~* '^#[a-f0-9]{6}$'),
    headbar varchar(7) CHECK (headbar is null or headbar ~* '^#[a-f0-9]{6}$'),
    yes varchar(7) CHECK (yes is null or yes ~* '^#[a-f0-9]{6}$'),
    no varchar(7) CHECK (no is null or no ~* '^#[a-f0-9]{6}$'),
    abstain varchar(7) CHECK (abstain is null or abstain ~* '^#[a-f0-9]{6}$'),
    organization_id integer GENERATED ALWAYS AS (1) STORED NOT NULL
);




CREATE TABLE committee_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    description text,
    external_id varchar(256),
    default_meeting_id integer UNIQUE,
    parent_id integer,
    organization_id integer GENERATED ALWAYS AS (1) STORED NOT NULL
);



comment on column committee_t.external_id is 'unique';


CREATE TABLE meeting_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    external_id varchar(256),
    welcome_title varchar(256) DEFAULT 'Welcome to OpenSlides',
    welcome_text text DEFAULT 'Space for your welcome text.',
    name varchar(100) NOT NULL DEFAULT 'OpenSlides',
    is_active_in_organization_id integer,
    is_archived_in_organization_id integer,
    description varchar(100) DEFAULT 'Presentation and assembly system',
    location varchar(256),
    start_time timestamptz,
    end_time timestamptz,
    locked_from_inside boolean,
    imported_at timestamptz,
    language varchar(256) CONSTRAINT enum_meeting_language CHECK (language IN ('en', 'de', 'it', 'es', 'ru', 'cs', 'fr')) DEFAULT 'en',
    jitsi_domain varchar(256),
    jitsi_room_name varchar(256),
    jitsi_room_password varchar(256),
    template_for_organization_id integer,
    enable_anonymous boolean DEFAULT False,
    custom_translations jsonb,
    conference_show boolean DEFAULT False,
    conference_auto_connect boolean DEFAULT False,
    conference_los_restriction boolean DEFAULT True,
    conference_stream_url varchar(256),
    conference_stream_poster_url varchar(256),
    conference_open_microphone boolean DEFAULT False,
    conference_open_video boolean DEFAULT False,
    conference_auto_connect_next_speakers integer DEFAULT 0,
    conference_enable_helpdesk boolean DEFAULT False,
    applause_enable boolean DEFAULT False,
    applause_type varchar(256) CONSTRAINT enum_meeting_applause_type CHECK (applause_type IN ('applause-type-bar', 'applause-type-particles')) DEFAULT 'applause-type-bar',
    applause_show_level boolean DEFAULT False,
    applause_min_amount integer CONSTRAINT minimum_applause_min_amount CHECK (applause_min_amount >= 0) DEFAULT 1,
    applause_max_amount integer CONSTRAINT minimum_applause_max_amount CHECK (applause_max_amount >= 0) DEFAULT 0,
    applause_timeout integer CONSTRAINT minimum_applause_timeout CHECK (applause_timeout >= 0) DEFAULT 5,
    applause_particle_image_url varchar(256),
    projector_countdown_default_time integer NOT NULL DEFAULT 60,
    projector_countdown_warning_time integer NOT NULL CONSTRAINT minimum_projector_countdown_warning_time CHECK (projector_countdown_warning_time >= 0) DEFAULT 0,
    export_csv_encoding varchar(256) CONSTRAINT enum_meeting_export_csv_encoding CHECK (export_csv_encoding IN ('utf-8', 'iso-8859-15')) DEFAULT 'utf-8',
    export_csv_separator varchar(256) DEFAULT ';',
    export_pdf_pagenumber_alignment varchar(256) CONSTRAINT enum_meeting_export_pdf_pagenumber_alignment CHECK (export_pdf_pagenumber_alignment IN ('left', 'right', 'center')) DEFAULT 'center',
    export_pdf_fontsize integer CONSTRAINT enum_meeting_export_pdf_fontsize CHECK (export_pdf_fontsize IN (10, 11, 12)) DEFAULT 10,
    export_pdf_line_height double precision CONSTRAINT minimum_export_pdf_line_height CHECK (export_pdf_line_height >= 1.0) DEFAULT 1.25,
    export_pdf_page_margin_left integer CONSTRAINT minimum_export_pdf_page_margin_left CHECK (export_pdf_page_margin_left >= 0) DEFAULT 20,
    export_pdf_page_margin_top integer CONSTRAINT minimum_export_pdf_page_margin_top CHECK (export_pdf_page_margin_top >= 0) DEFAULT 25,
    export_pdf_page_margin_right integer CONSTRAINT minimum_export_pdf_page_margin_right CHECK (export_pdf_page_margin_right >= 0) DEFAULT 20,
    export_pdf_page_margin_bottom integer CONSTRAINT minimum_export_pdf_page_margin_bottom CHECK (export_pdf_page_margin_bottom >= 0) DEFAULT 20,
    export_pdf_pagesize varchar(256) CONSTRAINT enum_meeting_export_pdf_pagesize CHECK (export_pdf_pagesize IN ('A4', 'A5')) DEFAULT 'A4',
    agenda_show_subtitles boolean DEFAULT False,
    agenda_enable_numbering boolean DEFAULT True,
    agenda_number_prefix varchar(20),
    agenda_numeral_system varchar(256) CONSTRAINT enum_meeting_agenda_numeral_system CHECK (agenda_numeral_system IN ('arabic', 'roman')) DEFAULT 'arabic',
    agenda_item_creation varchar(256) CONSTRAINT enum_meeting_agenda_item_creation CHECK (agenda_item_creation IN ('always', 'never', 'default_yes', 'default_no')) DEFAULT 'default_no',
    agenda_new_items_default_visibility varchar(256) CONSTRAINT enum_meeting_agenda_new_items_default_visibility CHECK (agenda_new_items_default_visibility IN ('common', 'internal', 'hidden')) DEFAULT 'internal',
    agenda_show_internal_items_on_projector boolean DEFAULT False,
    agenda_show_topic_navigation_on_detail_view boolean DEFAULT False,
    list_of_speakers_amount_last_on_projector integer CONSTRAINT minimum_list_of_speakers_amount_last_on_projector CHECK (list_of_speakers_amount_last_on_projector >= -1) DEFAULT 0,
    list_of_speakers_amount_next_on_projector integer CONSTRAINT minimum_list_of_speakers_amount_next_on_projector CHECK (list_of_speakers_amount_next_on_projector >= -1) DEFAULT -1,
    list_of_speakers_couple_countdown boolean DEFAULT True,
    list_of_speakers_show_amount_of_speakers_on_slide boolean DEFAULT True,
    list_of_speakers_present_users_only boolean DEFAULT False,
    list_of_speakers_show_first_contribution boolean DEFAULT False,
    list_of_speakers_hide_contribution_count boolean DEFAULT False,
    list_of_speakers_allow_multiple_speakers boolean DEFAULT False,
    list_of_speakers_enable_point_of_order_speakers boolean DEFAULT True,
    list_of_speakers_can_create_point_of_order_for_others boolean DEFAULT False,
    list_of_speakers_enable_point_of_order_categories boolean DEFAULT False,
    list_of_speakers_closing_disables_point_of_order boolean DEFAULT False,
    list_of_speakers_enable_pro_contra_speech boolean DEFAULT False,
    list_of_speakers_can_set_contribution_self boolean DEFAULT False,
    list_of_speakers_speaker_note_for_everyone boolean DEFAULT True,
    list_of_speakers_initially_closed boolean DEFAULT False,
    list_of_speakers_default_structure_level_time integer CONSTRAINT minimum_list_of_speakers_default_structure_level_time CHECK (list_of_speakers_default_structure_level_time >= 0),
    list_of_speakers_enable_interposed_question boolean,
    list_of_speakers_intervention_time integer,
    motions_default_workflow_id integer NOT NULL UNIQUE,
    motions_default_amendment_workflow_id integer NOT NULL UNIQUE,
    motions_preamble text DEFAULT 'The assembly may decide:',
    motions_default_line_numbering varchar(256) CONSTRAINT enum_meeting_motions_default_line_numbering CHECK (motions_default_line_numbering IN ('outside', 'inline', 'none')) DEFAULT 'outside',
    motions_line_length integer CONSTRAINT minimum_motions_line_length CHECK (motions_line_length >= 40) DEFAULT 85,
    motions_reason_required boolean DEFAULT False,
    motions_origin_motion_toggle_default boolean DEFAULT False,
    motions_enable_origin_motion_display boolean DEFAULT False,
    motions_enable_text_on_projector boolean DEFAULT True,
    motions_enable_reason_on_projector boolean DEFAULT False,
    motions_enable_sidebox_on_projector boolean DEFAULT False,
    motions_enable_recommendation_on_projector boolean DEFAULT True,
    motions_hide_metadata_background boolean DEFAULT False,
    motions_show_referring_motions boolean DEFAULT True,
    motions_show_sequential_number boolean DEFAULT True,
    motions_create_enable_additional_submitter_text boolean,
    motions_recommendations_by varchar(256),
    motions_block_slide_columns integer CONSTRAINT minimum_motions_block_slide_columns CHECK (motions_block_slide_columns >= 1),
    motions_recommendation_text_mode varchar(256) CONSTRAINT enum_meeting_motions_recommendation_text_mode CHECK (motions_recommendation_text_mode IN ('original', 'changed', 'diff', 'agreed')) DEFAULT 'diff',
    motions_default_sorting varchar(256) CONSTRAINT enum_meeting_motions_default_sorting CHECK (motions_default_sorting IN ('number', 'weight')) DEFAULT 'number',
    motions_number_type varchar(256) CONSTRAINT enum_meeting_motions_number_type CHECK (motions_number_type IN ('per_category', 'serially_numbered', 'manually')) DEFAULT 'per_category',
    motions_number_min_digits integer DEFAULT 2,
    motions_number_with_blank boolean DEFAULT False,
    motions_amendments_enabled boolean DEFAULT True,
    motions_amendments_in_main_list boolean DEFAULT True,
    motions_amendments_of_amendments boolean DEFAULT False,
    motions_amendments_prefix varchar(256) DEFAULT '-Ä',
    motions_amendments_text_mode varchar(256) CONSTRAINT enum_meeting_motions_amendments_text_mode CHECK (motions_amendments_text_mode IN ('freestyle', 'fulltext', 'paragraph')) DEFAULT 'paragraph',
    motions_amendments_multiple_paragraphs boolean DEFAULT True,
    motions_supporters_min_amount integer CONSTRAINT minimum_motions_supporters_min_amount CHECK (motions_supporters_min_amount >= 0) DEFAULT 0,
    motions_enable_editor boolean,
    motions_enable_working_group_speaker boolean,
    motions_export_title varchar(256) DEFAULT 'Motions',
    motions_export_preamble text,
    motions_export_submitter_recommendation boolean DEFAULT True,
    motions_export_follow_recommendation boolean DEFAULT False,
    motion_poll_ballot_paper_selection varchar(256) CONSTRAINT enum_meeting_motion_poll_ballot_paper_selection CHECK (motion_poll_ballot_paper_selection IN ('NUMBER_OF_DELEGATES', 'NUMBER_OF_ALL_PARTICIPANTS', 'CUSTOM_NUMBER')) DEFAULT 'CUSTOM_NUMBER',
    motion_poll_ballot_paper_number integer DEFAULT 8,
    motion_poll_default_type varchar(256) DEFAULT 'pseudoanonymous',
    motion_poll_default_method varchar(256) DEFAULT 'YNA',
    motion_poll_default_onehundred_percent_base varchar(256) CONSTRAINT enum_meeting_motion_poll_default_onehundred_percent_base CHECK (motion_poll_default_onehundred_percent_base IN ('Y', 'YN', 'YNA', 'N', 'valid', 'cast', 'entitled', 'entitled_present', 'disabled')) DEFAULT 'YNA',
    motion_poll_default_backend varchar(256) CONSTRAINT enum_meeting_motion_poll_default_backend CHECK (motion_poll_default_backend IN ('long', 'fast')) DEFAULT 'fast',
    motion_poll_projection_name_order_first varchar(256) NOT NULL CONSTRAINT enum_meeting_motion_poll_projection_name_order_first CHECK (motion_poll_projection_name_order_first IN ('first_name', 'last_name')) DEFAULT 'last_name',
    motion_poll_projection_max_columns integer NOT NULL DEFAULT 6,
    users_enable_presence_view boolean DEFAULT False,
    users_enable_vote_weight boolean DEFAULT False,
    users_allow_self_set_present boolean DEFAULT True,
    users_pdf_welcometitle varchar(256) DEFAULT 'Welcome to OpenSlides',
    users_pdf_welcometext text DEFAULT '[Place for your welcome and help text.]',
    users_pdf_wlan_ssid varchar(256),
    users_pdf_wlan_password varchar(256),
    users_pdf_wlan_encryption varchar(256) CONSTRAINT enum_meeting_users_pdf_wlan_encryption CHECK (users_pdf_wlan_encryption IN ('', 'WEP', 'WPA', 'nopass')) DEFAULT 'WPA',
    users_email_sender varchar(256) DEFAULT 'OpenSlides',
    users_email_replyto varchar(256),
    users_email_subject varchar(256) DEFAULT 'OpenSlides access data',
    users_email_body text DEFAULT 'Dear {name},

this is your personal OpenSlides login:

{url}
Username: {username}
Password: {password}


This email was generated automatically.',
    users_enable_vote_delegations boolean,
    users_forbid_delegator_in_list_of_speakers boolean,
    users_forbid_delegator_as_submitter boolean,
    users_forbid_delegator_as_supporter boolean,
    users_forbid_delegator_to_vote boolean,
    assignments_export_title varchar(256) DEFAULT 'Elections',
    assignments_export_preamble text,
    assignment_poll_ballot_paper_selection varchar(256) CONSTRAINT enum_meeting_assignment_poll_ballot_paper_selection CHECK (assignment_poll_ballot_paper_selection IN ('NUMBER_OF_DELEGATES', 'NUMBER_OF_ALL_PARTICIPANTS', 'CUSTOM_NUMBER')) DEFAULT 'CUSTOM_NUMBER',
    assignment_poll_ballot_paper_number integer DEFAULT 8,
    assignment_poll_add_candidates_to_list_of_speakers boolean DEFAULT False,
    assignment_poll_enable_max_votes_per_option boolean DEFAULT False,
    assignment_poll_sort_poll_result_by_votes boolean DEFAULT True,
    assignment_poll_default_type varchar(256) DEFAULT 'pseudoanonymous',
    assignment_poll_default_method varchar(256) DEFAULT 'Y',
    assignment_poll_default_onehundred_percent_base varchar(256) CONSTRAINT enum_meeting_assignment_poll_default_onehundred_percent_base CHECK (assignment_poll_default_onehundred_percent_base IN ('Y', 'YN', 'YNA', 'N', 'valid', 'cast', 'entitled', 'entitled_present', 'disabled')) DEFAULT 'valid',
    assignment_poll_default_backend varchar(256) CONSTRAINT enum_meeting_assignment_poll_default_backend CHECK (assignment_poll_default_backend IN ('long', 'fast')) DEFAULT 'fast',
    poll_ballot_paper_selection varchar(256) CONSTRAINT enum_meeting_poll_ballot_paper_selection CHECK (poll_ballot_paper_selection IN ('NUMBER_OF_DELEGATES', 'NUMBER_OF_ALL_PARTICIPANTS', 'CUSTOM_NUMBER')),
    poll_ballot_paper_number integer,
    poll_sort_poll_result_by_votes boolean,
    poll_default_type varchar(256) DEFAULT 'analog',
    poll_default_method varchar(256),
    poll_default_onehundred_percent_base varchar(256) CONSTRAINT enum_meeting_poll_default_onehundred_percent_base CHECK (poll_default_onehundred_percent_base IN ('Y', 'YN', 'YNA', 'N', 'valid', 'cast', 'entitled', 'entitled_present', 'disabled')) DEFAULT 'YNA',
    poll_default_backend varchar(256) CONSTRAINT enum_meeting_poll_default_backend CHECK (poll_default_backend IN ('long', 'fast')) DEFAULT 'fast',
    poll_default_live_voting_enabled boolean DEFAULT False,
    poll_couple_countdown boolean DEFAULT True,
    logo_projector_main_id integer UNIQUE,
    logo_projector_header_id integer UNIQUE,
    logo_web_header_id integer UNIQUE,
    logo_pdf_header_l_id integer UNIQUE,
    logo_pdf_header_r_id integer UNIQUE,
    logo_pdf_footer_l_id integer UNIQUE,
    logo_pdf_footer_r_id integer UNIQUE,
    logo_pdf_ballot_paper_id integer UNIQUE,
    font_regular_id integer UNIQUE,
    font_italic_id integer UNIQUE,
    font_bold_id integer UNIQUE,
    font_bold_italic_id integer UNIQUE,
    font_monospace_id integer UNIQUE,
    font_chyron_speaker_name_id integer UNIQUE,
    font_projector_h1_id integer UNIQUE,
    font_projector_h2_id integer UNIQUE,
    committee_id integer NOT NULL,
    reference_projector_id integer NOT NULL UNIQUE,
    list_of_speakers_countdown_id integer UNIQUE,
    poll_countdown_id integer UNIQUE,
    default_group_id integer NOT NULL UNIQUE,
    admin_group_id integer UNIQUE,
    anonymous_group_id integer UNIQUE
);



comment on column meeting_t.external_id is 'unique in committee';
comment on column meeting_t.is_active_in_organization_id is 'Backrelation and boolean flag at once';
comment on column meeting_t.is_archived_in_organization_id is 'Backrelation and boolean flag at once';
comment on column meeting_t.list_of_speakers_default_structure_level_time is '0 disables structure level countdowns.';
comment on column meeting_t.list_of_speakers_intervention_time is '0 disables intervention speakers.';
comment on column meeting_t.poll_default_live_voting_enabled is 'Defines default ''poll.live_voting_enabled'' option suggested to user. Is not used in the validations.';


CREATE TABLE structure_level_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    color varchar(7) CHECK (color is null or color ~* '^#[a-f0-9]{6}$'),
    default_time integer CONSTRAINT minimum_default_time CHECK (default_time >= 0),
    meeting_id integer NOT NULL
);




CREATE TABLE group_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    external_id varchar(256),
    name varchar(256) NOT NULL,
    permissions varchar(256)[] CONSTRAINT enum_group_permissions CHECK (permissions <@ ARRAY['agenda_item.can_manage', 'agenda_item.can_see', 'agenda_item.can_see_internal', 'assignment.can_manage', 'assignment.can_nominate_other', 'assignment.can_nominate_self', 'assignment.can_see', 'chat.can_manage', 'list_of_speakers.can_be_speaker', 'list_of_speakers.can_manage', 'list_of_speakers.can_see', 'list_of_speakers.can_manage_moderator_notes', 'list_of_speakers.can_see_moderator_notes', 'mediafile.can_manage', 'mediafile.can_see', 'meeting.can_manage_logos_and_fonts', 'meeting.can_manage_settings', 'meeting.can_see_autopilot', 'meeting.can_see_frontpage', 'meeting.can_see_history', 'meeting.can_see_livestream', 'motion.can_create', 'motion.can_create_amendments', 'motion.can_forward', 'motion.can_manage', 'motion.can_manage_metadata', 'motion.can_manage_polls', 'motion.can_see', 'motion.can_see_internal', 'motion.can_see_origin', 'motion.can_support', 'poll.can_manage', 'poll.can_see_progress', 'projector.can_manage', 'projector.can_see', 'tag.can_manage', 'user.can_manage', 'user.can_manage_presence', 'user.can_see_sensitive_data', 'user.can_see', 'user.can_update', 'user.can_edit_own_delegation']::varchar[]),
    weight integer,
    used_as_motion_poll_default_id integer,
    used_as_assignment_poll_default_id integer,
    used_as_topic_poll_default_id integer,
    used_as_poll_default_id integer,
    meeting_id integer NOT NULL
);



comment on column group_t.external_id is 'unique in meeting';


CREATE TABLE personal_note_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    note text,
    star boolean,
    meeting_user_id integer NOT NULL,
    content_object_id varchar(100),
    content_object_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    CONSTRAINT valid_content_object_id_part1 CHECK (split_part(content_object_id, '/', 1) IN ('motion')),
    meeting_id integer NOT NULL
);




CREATE TABLE tag_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE agenda_item_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    item_number varchar(256),
    comment varchar(256),
    closed boolean DEFAULT False,
    type varchar(256) CONSTRAINT enum_agenda_item_type CHECK (type IN ('common', 'internal', 'hidden')) DEFAULT 'common',
    duration integer CONSTRAINT minimum_duration CHECK (duration >= 0),
    is_internal boolean,
    is_hidden boolean,
    level integer,
    weight integer,
    content_object_id varchar(100) NOT NULL,
    content_object_id_motion_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_motion_block_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion_block' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_assignment_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'assignment' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_topic_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'topic' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    CONSTRAINT valid_content_object_id_part1 CHECK (split_part(content_object_id, '/', 1) IN ('motion','motion_block','assignment','topic')),
    parent_id integer,
    meeting_id integer NOT NULL
);



comment on column agenda_item_t.duration is 'Given in seconds';
comment on column agenda_item_t.is_internal is 'Calculated by the server';
comment on column agenda_item_t.is_hidden is 'Calculated by the server';
comment on column agenda_item_t.level is 'Calculated by the server';


CREATE TABLE list_of_speakers_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    closed boolean DEFAULT False,
    sequential_number integer NOT NULL,
    moderator_notes text,
    content_object_id varchar(100) NOT NULL,
    content_object_id_motion_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_motion_block_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion_block' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_assignment_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'assignment' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_topic_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'topic' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_meeting_mediafile_id integer UNIQUE GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'meeting_mediafile' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    CONSTRAINT valid_content_object_id_part1 CHECK (split_part(content_object_id, '/', 1) IN ('motion','motion_block','assignment','topic','meeting_mediafile')),
    meeting_id integer NOT NULL
);



comment on column list_of_speakers_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE structure_level_list_of_speakers_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    structure_level_id integer NOT NULL,
    list_of_speakers_id integer NOT NULL,
    initial_time integer NOT NULL CONSTRAINT minimum_initial_time CHECK (initial_time >= 1),
    additional_time double precision,
    remaining_time double precision NOT NULL,
    current_start_time timestamptz,
    meeting_id integer NOT NULL
);



comment on column structure_level_list_of_speakers_t.initial_time is 'The initial time of this structure_level for this LoS';
comment on column structure_level_list_of_speakers_t.additional_time is 'The summed added time of this structure_level for this LoS';
comment on column structure_level_list_of_speakers_t.remaining_time is 'The currently remaining time of this structure_level for this LoS';
comment on column structure_level_list_of_speakers_t.current_start_time is 'The current start time of a speaker for this structure_level. Is only set if a currently speaking speaker exists';


CREATE TABLE point_of_order_category_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    text varchar(256) NOT NULL,
    rank integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE speaker_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    begin_time timestamptz,
    end_time timestamptz,
    pause_time timestamptz,
    unpause_time timestamptz,
    total_pause integer,
    weight integer DEFAULT 10000,
    speech_state varchar(256) CONSTRAINT enum_speaker_speech_state CHECK (speech_state IN ('contribution', 'pro', 'contra', 'intervention', 'interposed_question')),
    note varchar(250),
    point_of_order boolean,
    list_of_speakers_id integer NOT NULL,
    structure_level_list_of_speakers_id integer,
    meeting_user_id integer,
    point_of_order_category_id integer,
    meeting_id integer NOT NULL
);




CREATE TABLE topic_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title varchar(256) NOT NULL,
    text text,
    sequential_number integer NOT NULL,
    meeting_id integer NOT NULL
);



comment on column topic_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE motion_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    number varchar(256),
    number_value integer,
    sequential_number integer NOT NULL,
    title varchar(256) NOT NULL,
    text text,
    text_hash varchar(256),
    amendment_paragraphs jsonb,
    modified_final_version text,
    reason text,
    category_weight integer DEFAULT 10000,
    state_extension varchar(256),
    recommendation_extension varchar(256),
    sort_weight integer DEFAULT 10000,
    created timestamptz,
    last_modified timestamptz,
    workflow_timestamp timestamptz,
    start_line_number integer CONSTRAINT minimum_start_line_number CHECK (start_line_number >= 1) DEFAULT 1,
    forwarded timestamptz,
    additional_submitter varchar(256),
    marked_forwarded boolean,
    lead_motion_id integer,
    sort_parent_id integer,
    origin_id integer,
    origin_meeting_id integer,
    state_id integer NOT NULL,
    recommendation_id integer,
    category_id integer,
    block_id integer,
    meeting_id integer NOT NULL
);



comment on column motion_t.number_value is 'The number value of this motion. This number is auto-generated and read-only.';
comment on column motion_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';
comment on column motion_t.marked_forwarded is 'Forwarded amendments can be marked as such. This is just optional, however. Forwarded amendments can also have this field set to false.';


CREATE TABLE motion_submitter_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    weight integer,
    meeting_user_id integer NOT NULL,
    motion_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE motion_editor_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    weight integer,
    meeting_user_id integer NOT NULL,
    motion_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE motion_working_group_speaker_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    weight integer,
    meeting_user_id integer NOT NULL,
    motion_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE motion_comment_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    comment text,
    motion_id integer NOT NULL,
    section_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE motion_comment_section_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    weight integer DEFAULT 10000,
    sequential_number integer NOT NULL,
    submitter_can_write boolean,
    meeting_id integer NOT NULL
);



comment on column motion_comment_section_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE motion_category_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    prefix varchar(256),
    weight integer DEFAULT 10000,
    level integer,
    sequential_number integer NOT NULL,
    parent_id integer,
    meeting_id integer NOT NULL
);



comment on column motion_category_t.level is 'Calculated field.';
comment on column motion_category_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE motion_block_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title varchar(256) NOT NULL,
    internal boolean,
    sequential_number integer NOT NULL,
    meeting_id integer NOT NULL
);



comment on column motion_block_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE motion_change_recommendation_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    rejected boolean DEFAULT False,
    internal boolean DEFAULT False,
    type varchar(256) CONSTRAINT enum_motion_change_recommendation_type CHECK (type IN ('replacement', 'insertion', 'deletion', 'other')) DEFAULT 'replacement',
    other_description varchar(256),
    line_from integer CONSTRAINT minimum_line_from CHECK (line_from >= 0),
    line_to integer CONSTRAINT minimum_line_to CHECK (line_to >= 0),
    text text,
    creation_time timestamptz,
    motion_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE motion_state_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    weight integer NOT NULL,
    recommendation_label varchar(256),
    is_internal boolean,
    css_class varchar(256) NOT NULL CONSTRAINT enum_motion_state_css_class CHECK (css_class IN ('grey', 'red', 'green', 'lightblue', 'yellow')) DEFAULT 'lightblue',
    restrictions varchar(256)[] CONSTRAINT enum_motion_state_restrictions CHECK (restrictions <@ ARRAY['motion.can_see_internal', 'motion.can_manage_metadata', 'motion.can_manage', 'is_submitter']::varchar[]) DEFAULT '{}',
    allow_support boolean DEFAULT False,
    allow_create_poll boolean DEFAULT False,
    allow_submitter_edit boolean DEFAULT False,
    set_number boolean DEFAULT True,
    show_state_extension_field boolean DEFAULT False,
    show_recommendation_extension_field boolean DEFAULT False,
    merge_amendment_into_final varchar(256) CONSTRAINT enum_motion_state_merge_amendment_into_final CHECK (merge_amendment_into_final IN ('do_not_merge', 'undefined', 'do_merge')) DEFAULT 'undefined',
    allow_motion_forwarding boolean DEFAULT False,
    allow_amendment_forwarding boolean,
    set_workflow_timestamp boolean DEFAULT False,
    submitter_withdraw_state_id integer,
    workflow_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE motion_workflow_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    sequential_number integer NOT NULL,
    first_state_id integer NOT NULL UNIQUE,
    meeting_id integer NOT NULL
);



comment on column motion_workflow_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE poll_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title varchar(256) NOT NULL,
    description text,
    method varchar(256) NOT NULL CONSTRAINT enum_poll_method CHECK (method IN ('motion', 'selection', 'rating', 'single_transferable_vote')),
    config jsonb,
    visibility varchar(256) CONSTRAINT enum_poll_visibility CHECK (visibility IN ('manually', 'named', 'open', 'secret')),
    state varchar(256) CONSTRAINT enum_poll_state CHECK (state IN ('created', 'started', 'finished', 'published')) DEFAULT 'created',
    result jsonb,
    sequential_number integer NOT NULL,
    content_object_id varchar(100) NOT NULL,
    content_object_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_assignment_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'assignment' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_topic_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'topic' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    CONSTRAINT valid_content_object_id_part1 CHECK (split_part(content_object_id, '/', 1) IN ('motion','assignment','topic')),
    meeting_id integer NOT NULL
);



comment on column poll_t.config is 'Values to configure the poll. Depends on the value in poll/method.';
comment on column poll_t.result is 'Calculated result. The format depends on the value in poll/method';
comment on column poll_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE vote_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    weight decimal(16,6) DEFAULT '1.000000',
    value jsonb,
    poll_id integer NOT NULL,
    acting_user_id integer,
    represented_user_id integer,
    meeting_id integer NOT NULL
);




CREATE TABLE assignment_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title varchar(256) NOT NULL,
    description text,
    open_posts integer CONSTRAINT minimum_open_posts CHECK (open_posts >= 0) DEFAULT 0,
    phase varchar(256) CONSTRAINT enum_assignment_phase CHECK (phase IN ('search', 'voting', 'finished')) DEFAULT 'search',
    default_poll_description text,
    number_poll_candidates boolean,
    sequential_number integer NOT NULL,
    meeting_id integer NOT NULL
);



comment on column assignment_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE assignment_candidate_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    weight integer DEFAULT 10000,
    assignment_id integer NOT NULL,
    meeting_user_id integer,
    meeting_id integer NOT NULL
);




CREATE TABLE poll_candidate_list_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE poll_candidate_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    poll_candidate_list_id integer NOT NULL,
    user_id integer,
    weight integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE mediafile_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title varchar(256),
    is_directory boolean,
    filesize integer,
    filename varchar(256),
    mimetype varchar(256),
    pdf_information jsonb,
    create_timestamp timestamptz,
    token varchar(256),
    published_to_meetings_in_organization_id integer,
    parent_id integer,
    owner_id varchar(100) NOT NULL,
    owner_id_meeting_id integer GENERATED ALWAYS AS (CASE WHEN split_part(owner_id, '/', 1) = 'meeting' THEN cast(split_part(owner_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    owner_id_organization_id integer GENERATED ALWAYS AS (CASE WHEN split_part(owner_id, '/', 1) = 'organization' THEN cast(split_part(owner_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    CONSTRAINT valid_owner_id_part1 CHECK (split_part(owner_id, '/', 1) IN ('meeting','organization'))
);



comment on column mediafile_t.title is 'Title and parent_id must be unique.';
comment on column mediafile_t.filesize is 'In bytes, not the human readable format anymore.';
comment on column mediafile_t.filename is 'The uploaded filename. Will be used for downloading. Only writeable on create.';


CREATE TABLE meeting_mediafile_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    mediafile_id integer NOT NULL,
    meeting_id integer NOT NULL,
    is_public boolean NOT NULL
);



comment on column meeting_mediafile_t.is_public is 'Calculated in actions. Used to discern whether the (meeting-)mediafile can be seen by everyone, because, in the case of inherited_access_group_ids == [], it would otherwise not be clear. inherited_access_group_ids == [] can have two causes: cancelling access groups (=> is_public := false) or no access groups at all (=> is_public := true)';


CREATE TABLE projector_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256),
    is_internal boolean DEFAULT False,
    scale integer DEFAULT 0,
    scroll integer CONSTRAINT minimum_scroll CHECK (scroll >= 0) DEFAULT 0,
    width integer CONSTRAINT minimum_width CHECK (width >= 1) DEFAULT 1200,
    aspect_ratio_numerator integer CONSTRAINT minimum_aspect_ratio_numerator CHECK (aspect_ratio_numerator >= 1) DEFAULT 16,
    aspect_ratio_denominator integer CONSTRAINT minimum_aspect_ratio_denominator CHECK (aspect_ratio_denominator >= 1) DEFAULT 9,
    color varchar(7) CHECK (color is null or color ~* '^#[a-f0-9]{6}$') DEFAULT '#000000',
    background_color varchar(7) CHECK (background_color is null or background_color ~* '^#[a-f0-9]{6}$') DEFAULT '#ffffff',
    header_background_color varchar(7) CHECK (header_background_color is null or header_background_color ~* '^#[a-f0-9]{6}$') DEFAULT '#317796',
    header_font_color varchar(7) CHECK (header_font_color is null or header_font_color ~* '^#[a-f0-9]{6}$') DEFAULT '#f5f5f5',
    header_h1_color varchar(7) CHECK (header_h1_color is null or header_h1_color ~* '^#[a-f0-9]{6}$') DEFAULT '#317796',
    chyron_background_color varchar(7) CHECK (chyron_background_color is null or chyron_background_color ~* '^#[a-f0-9]{6}$') DEFAULT '#317796',
    chyron_background_color_2 varchar(7) CHECK (chyron_background_color_2 is null or chyron_background_color_2 ~* '^#[a-f0-9]{6}$') DEFAULT '#134768',
    chyron_font_color varchar(7) CHECK (chyron_font_color is null or chyron_font_color ~* '^#[a-f0-9]{6}$') DEFAULT '#ffffff',
    chyron_font_color_2 varchar(7) CHECK (chyron_font_color_2 is null or chyron_font_color_2 ~* '^#[a-f0-9]{6}$') DEFAULT '#ffffff',
    show_header_footer boolean DEFAULT True,
    show_title boolean DEFAULT True,
    show_logo boolean DEFAULT True,
    show_clock boolean DEFAULT True,
    sequential_number integer NOT NULL,
    used_as_default_projector_for_agenda_item_list_in_meeting_id integer,
    used_as_default_projector_for_topic_in_meeting_id integer,
    used_as_default_projector_for_list_of_speakers_in_meeting_id integer,
    used_as_default_projector_for_current_los_in_meeting_id integer,
    used_as_default_projector_for_motion_in_meeting_id integer,
    used_as_default_projector_for_amendment_in_meeting_id integer,
    used_as_default_projector_for_motion_block_in_meeting_id integer,
    used_as_default_projector_for_assignment_in_meeting_id integer,
    used_as_default_projector_for_mediafile_in_meeting_id integer,
    used_as_default_projector_for_message_in_meeting_id integer,
    used_as_default_projector_for_countdown_in_meeting_id integer,
    used_as_default_projector_for_assignment_poll_in_meeting_id integer,
    used_as_default_projector_for_motion_poll_in_meeting_id integer,
    used_as_default_projector_for_poll_in_meeting_id integer,
    meeting_id integer NOT NULL
);



comment on column projector_t.sequential_number is 'The (positive) serial number of this model in its meeting. This number is auto-generated and read-only.';


CREATE TABLE projection_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    options jsonb,
    stable boolean DEFAULT False,
    weight integer,
    type varchar(256),
    current_projector_id integer,
    preview_projector_id integer,
    history_projector_id integer,
    content_object_id varchar(100) NOT NULL,
    content_object_id_meeting_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'meeting' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_meeting_mediafile_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'meeting_mediafile' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_list_of_speakers_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'list_of_speakers' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_motion_block_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'motion_block' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_assignment_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'assignment' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_agenda_item_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'agenda_item' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_topic_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'topic' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_poll_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'poll' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_projector_message_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'projector_message' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    content_object_id_projector_countdown_id integer GENERATED ALWAYS AS (CASE WHEN split_part(content_object_id, '/', 1) = 'projector_countdown' THEN cast(split_part(content_object_id, '/', 2) AS INTEGER) ELSE null END) STORED,
    CONSTRAINT valid_content_object_id_part1 CHECK (split_part(content_object_id, '/', 1) IN ('meeting','motion','meeting_mediafile','list_of_speakers','motion_block','assignment','agenda_item','topic','poll','projector_message','projector_countdown')),
    meeting_id integer NOT NULL
);



/*
 Fields without SQL definition for table projection

    projection/content: type:JSON is marked as a calculated field and not generated in schema

*/

CREATE TABLE projector_message_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    message text,
    meeting_id integer NOT NULL
);




CREATE TABLE projector_countdown_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title varchar(256) NOT NULL,
    description varchar(256) DEFAULT '',
    default_time integer,
    countdown_time double precision DEFAULT 60,
    running boolean DEFAULT False,
    meeting_id integer NOT NULL
);




CREATE TABLE chat_group_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    weight integer DEFAULT 10000,
    meeting_id integer NOT NULL
);




CREATE TABLE chat_message_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    content text NOT NULL,
    created timestamptz NOT NULL,
    meeting_user_id integer,
    chat_group_id integer NOT NULL,
    meeting_id integer NOT NULL
);




CREATE TABLE action_worker_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL,
    state varchar(256) NOT NULL CONSTRAINT enum_action_worker_state CHECK (state IN ('running', 'end', 'aborted')),
    created timestamptz NOT NULL,
    timestamp timestamptz NOT NULL,
    result jsonb,
    user_id integer NOT NULL
);



comment on column action_worker_t.user_id is 'Id of the calling user. If the action is called via internal route, the value will be -1.';


CREATE TABLE import_preview_t (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name varchar(256) NOT NULL CONSTRAINT enum_import_preview_name CHECK (name IN ('account', 'participant', 'topic', 'committee', 'motion')),
    state varchar(256) NOT NULL CONSTRAINT enum_import_preview_state CHECK (state IN ('warning', 'error', 'done')),
    created timestamptz NOT NULL,
    result jsonb
);





-- Intermediate table definitions

CREATE TABLE nm_meeting_user_supported_motion_ids_motion_t (
    meeting_user_id integer NOT NULL REFERENCES meeting_user_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    motion_id integer NOT NULL REFERENCES motion_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (meeting_user_id, motion_id)
);

CREATE TABLE nm_meeting_user_structure_level_ids_structure_level_t (
    meeting_user_id integer NOT NULL REFERENCES meeting_user_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    structure_level_id integer NOT NULL REFERENCES structure_level_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (meeting_user_id, structure_level_id)
);

CREATE TABLE gm_organization_tag_tagged_ids_t (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    organization_tag_id integer NOT NULL REFERENCES organization_tag_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    tagged_id varchar(100) NOT NULL,
    tagged_id_committee_id integer GENERATED ALWAYS AS (CASE WHEN split_part(tagged_id, '/', 1) = 'committee' THEN cast(split_part(tagged_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES committee_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    tagged_id_meeting_id integer GENERATED ALWAYS AS (CASE WHEN split_part(tagged_id, '/', 1) = 'meeting' THEN cast(split_part(tagged_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES meeting_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    CONSTRAINT valid_tagged_id_part1 CHECK (split_part(tagged_id, '/', 1) IN ('committee', 'meeting')),
    CONSTRAINT unique_$organization_tag_id_$tagged_id UNIQUE (organization_tag_id, tagged_id)
);

CREATE TABLE nm_committee_manager_ids_user_t (
    committee_id integer NOT NULL REFERENCES committee_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    user_id integer NOT NULL REFERENCES user_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (committee_id, user_id)
);

CREATE TABLE nm_committee_all_child_ids_committee_t (
    all_child_id integer NOT NULL REFERENCES committee_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    all_parent_id integer NOT NULL REFERENCES committee_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (all_child_id, all_parent_id)
);

CREATE TABLE nm_committee_forward_to_committee_ids_committee_t (
    forward_to_committee_id integer NOT NULL REFERENCES committee_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    receive_forwardings_from_committee_id integer NOT NULL REFERENCES committee_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (forward_to_committee_id, receive_forwardings_from_committee_id)
);

CREATE TABLE nm_meeting_present_user_ids_user_t (
    meeting_id integer NOT NULL REFERENCES meeting_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    user_id integer NOT NULL REFERENCES user_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (meeting_id, user_id)
);

CREATE TABLE nm_group_meeting_user_ids_meeting_user_t (
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    meeting_user_id integer NOT NULL REFERENCES meeting_user_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (group_id, meeting_user_id)
);

CREATE TABLE nm_group_mmagi_meeting_mediafile_t (
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    meeting_mediafile_id integer NOT NULL REFERENCES meeting_mediafile_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (group_id, meeting_mediafile_id)
);

CREATE TABLE nm_group_mmiagi_meeting_mediafile_t (
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    meeting_mediafile_id integer NOT NULL REFERENCES meeting_mediafile_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (group_id, meeting_mediafile_id)
);

CREATE TABLE nm_group_read_comment_section_ids_motion_comment_section_t (
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    motion_comment_section_id integer NOT NULL REFERENCES motion_comment_section_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (group_id, motion_comment_section_id)
);

CREATE TABLE nm_group_write_comment_section_ids_motion_comment_section_t (
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    motion_comment_section_id integer NOT NULL REFERENCES motion_comment_section_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (group_id, motion_comment_section_id)
);

CREATE TABLE nm_group_poll_ids_poll_t (
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    poll_id integer NOT NULL REFERENCES poll_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (group_id, poll_id)
);

CREATE TABLE gm_tag_tagged_ids_t (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    tag_id integer NOT NULL REFERENCES tag_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    tagged_id varchar(100) NOT NULL,
    tagged_id_agenda_item_id integer GENERATED ALWAYS AS (CASE WHEN split_part(tagged_id, '/', 1) = 'agenda_item' THEN cast(split_part(tagged_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES agenda_item_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    tagged_id_assignment_id integer GENERATED ALWAYS AS (CASE WHEN split_part(tagged_id, '/', 1) = 'assignment' THEN cast(split_part(tagged_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES assignment_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    tagged_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(tagged_id, '/', 1) = 'motion' THEN cast(split_part(tagged_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES motion_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    CONSTRAINT valid_tagged_id_part1 CHECK (split_part(tagged_id, '/', 1) IN ('agenda_item', 'assignment', 'motion')),
    CONSTRAINT unique_$tag_id_$tagged_id UNIQUE (tag_id, tagged_id)
);

CREATE TABLE nm_motion_all_derived_motion_ids_motion_t (
    all_derived_motion_id integer NOT NULL REFERENCES motion_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    all_origin_id integer NOT NULL REFERENCES motion_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (all_derived_motion_id, all_origin_id)
);

CREATE TABLE nm_motion_identical_motion_ids_motion_t (
    identical_motion_id_1 integer NOT NULL REFERENCES motion_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    identical_motion_id_2 integer NOT NULL REFERENCES motion_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (identical_motion_id_1, identical_motion_id_2)
);

CREATE TABLE gm_motion_state_extension_reference_ids_t (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    motion_id integer NOT NULL REFERENCES motion_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    state_extension_reference_id varchar(100) NOT NULL,
    state_extension_reference_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(state_extension_reference_id, '/', 1) = 'motion' THEN cast(split_part(state_extension_reference_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES motion_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    CONSTRAINT valid_state_extension_reference_id_part1 CHECK (split_part(state_extension_reference_id, '/', 1) IN ('motion')),
    CONSTRAINT unique_$motion_id_$state_extension_reference_id UNIQUE (motion_id, state_extension_reference_id)
);

CREATE TABLE gm_motion_recommendation_extension_reference_ids_t (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    motion_id integer NOT NULL REFERENCES motion_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    recommendation_extension_reference_id varchar(100) NOT NULL,
    recommendation_extension_reference_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(recommendation_extension_reference_id, '/', 1) = 'motion' THEN cast(split_part(recommendation_extension_reference_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES motion_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    CONSTRAINT valid_recommendation_extension_reference_id_part1 CHECK (split_part(recommendation_extension_reference_id, '/', 1) IN ('motion')),
    CONSTRAINT unique_$motion_id_$recommendation_extension_reference_id UNIQUE (motion_id, recommendation_extension_reference_id)
);

CREATE TABLE nm_motion_state_next_state_ids_motion_state_t (
    next_state_id integer NOT NULL REFERENCES motion_state_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    previous_state_id integer NOT NULL REFERENCES motion_state_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (next_state_id, previous_state_id)
);

CREATE TABLE nm_poll_voted_ids_user_t (
    poll_id integer NOT NULL REFERENCES poll_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    user_id integer NOT NULL REFERENCES user_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (poll_id, user_id)
);

CREATE TABLE gm_meeting_mediafile_attachment_ids_t (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    meeting_mediafile_id integer NOT NULL REFERENCES meeting_mediafile_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    attachment_id varchar(100) NOT NULL,
    attachment_id_motion_id integer GENERATED ALWAYS AS (CASE WHEN split_part(attachment_id, '/', 1) = 'motion' THEN cast(split_part(attachment_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES motion_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    attachment_id_topic_id integer GENERATED ALWAYS AS (CASE WHEN split_part(attachment_id, '/', 1) = 'topic' THEN cast(split_part(attachment_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES topic_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    attachment_id_assignment_id integer GENERATED ALWAYS AS (CASE WHEN split_part(attachment_id, '/', 1) = 'assignment' THEN cast(split_part(attachment_id, '/', 2) AS INTEGER) ELSE null END) STORED REFERENCES assignment_t(id) ON DELETE CASCADE INITIALLY DEFERRED,
    CONSTRAINT valid_attachment_id_part1 CHECK (split_part(attachment_id, '/', 1) IN ('motion', 'topic', 'assignment')),
    CONSTRAINT unique_$meeting_mediafile_id_$attachment_id UNIQUE (meeting_mediafile_id, attachment_id)
);

CREATE TABLE nm_chat_group_read_group_ids_group_t (
    chat_group_id integer NOT NULL REFERENCES chat_group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (chat_group_id, group_id)
);

CREATE TABLE nm_chat_group_write_group_ids_group_t (
    chat_group_id integer NOT NULL REFERENCES chat_group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    group_id integer NOT NULL REFERENCES group_t (id) ON DELETE CASCADE INITIALLY DEFERRED,
    PRIMARY KEY (chat_group_id, group_id)
);


-- View definitions

CREATE VIEW "organization" AS SELECT *,
(select array_agg(g.id ORDER BY g.id) from gender_t g where g.organization_id = o.id) as gender_ids,
(select array_agg(c.id ORDER BY c.id) from committee_t c where c.organization_id = o.id) as committee_ids,
(select array_agg(m.id ORDER BY m.id) from meeting_t m where m.is_active_in_organization_id = o.id) as active_meeting_ids,
(select array_agg(m.id ORDER BY m.id) from meeting_t m where m.is_archived_in_organization_id = o.id) as archived_meeting_ids,
(select array_agg(m.id ORDER BY m.id) from meeting_t m where m.template_for_organization_id = o.id) as template_meeting_ids,
(select array_agg(ot.id ORDER BY ot.id) from organization_tag_t ot where ot.organization_id = o.id) as organization_tag_ids,
(select array_agg(t.id ORDER BY t.id) from theme_t t where t.organization_id = o.id) as theme_ids,
(select array_agg(m.id ORDER BY m.id) from mediafile_t m where m.owner_id_organization_id = o.id) as mediafile_ids,
(select array_agg(m.id ORDER BY m.id) from mediafile_t m where m.published_to_meetings_in_organization_id = o.id) as published_mediafile_ids,
(select array_agg(u.id ORDER BY u.id) from user_t u where u.organization_id = o.id) as user_ids
FROM organization_t o;


CREATE VIEW "user" AS SELECT *,
(select array_agg(n.meeting_id ORDER BY n.meeting_id) from nm_meeting_present_user_ids_user_t n where n.user_id = u.id) as is_present_in_meeting_ids,
(
  SELECT array_agg(DISTINCT committee_id ORDER BY committee_id)
  FROM (
    -- Select committee_ids from meetings the user is part of
    SELECT m.committee_id
    FROM meeting_user_t AS mu
    INNER JOIN nm_group_meeting_user_ids_meeting_user_t AS gmu ON mu.id = gmu.meeting_user_id
    INNER JOIN meeting_t AS m ON m.id = mu.meeting_id
    WHERE mu.user_id = u.id

    UNION

    -- Select committee_ids from committee managers
    SELECT cmu.committee_id
    FROM nm_committee_manager_ids_user_t cmu
    WHERE cmu.user_id = u.id

    UNION

    -- Select home_committee_id from user
    SELECT u.home_committee_id
    WHERE u.home_committee_id IS NOT NULL
  ) _
) AS committee_ids
,
(select array_agg(n.committee_id ORDER BY n.committee_id) from nm_committee_manager_ids_user_t n where n.user_id = u.id) as committee_management_ids,
(select array_agg(m.id ORDER BY m.id) from meeting_user_t m where m.user_id = u.id) as meeting_user_ids,
(select array_agg(n.poll_id ORDER BY n.poll_id) from nm_poll_voted_ids_user_t n where n.user_id = u.id) as poll_voted_ids,
(select array_agg(v.id ORDER BY v.id) from vote_t v where v.acting_user_id = u.id) as acting_vote_ids,
(select array_agg(v.id ORDER BY v.id) from vote_t v where v.represented_user_id = u.id) as represented_vote_ids,
(select array_agg(p.id ORDER BY p.id) from poll_candidate_t p where p.user_id = u.id) as poll_candidate_ids,
(
  SELECT array_agg(DISTINCT mu.meeting_id ORDER BY mu.meeting_id)
  FROM meeting_user_t mu
  INNER JOIN nm_group_meeting_user_ids_meeting_user_t AS gmu ON mu.id = gmu.meeting_user_id
  WHERE mu.user_id = u.id
) AS meeting_ids

FROM user_t u;

comment on column "user".committee_ids is 'Calculated field: Returns committee_ids, where the user is manager or member in a meeting';
comment on column "user".meeting_ids is 'Calculated. All ids from meetings calculated via meeting_user and group_ids as integers.';

CREATE VIEW "meeting_user" AS SELECT *,
(select array_agg(p.id ORDER BY p.id) from personal_note_t p where p.meeting_user_id = m.id) as personal_note_ids,
(select array_agg(s.id ORDER BY s.id) from speaker_t s where s.meeting_user_id = m.id) as speaker_ids,
(select array_agg(n.motion_id ORDER BY n.motion_id) from nm_meeting_user_supported_motion_ids_motion_t n where n.meeting_user_id = m.id) as supported_motion_ids,
(select array_agg(me.id ORDER BY me.id) from motion_editor_t me where me.meeting_user_id = m.id) as motion_editor_ids,
(select array_agg(mw.id ORDER BY mw.id) from motion_working_group_speaker_t mw where mw.meeting_user_id = m.id) as motion_working_group_speaker_ids,
(select array_agg(ms.id ORDER BY ms.id) from motion_submitter_t ms where ms.meeting_user_id = m.id) as motion_submitter_ids,
(select array_agg(a.id ORDER BY a.id) from assignment_candidate_t a where a.meeting_user_id = m.id) as assignment_candidate_ids,
(select array_agg(mu.id ORDER BY mu.id) from meeting_user_t mu where mu.vote_delegated_to_id = m.id) as vote_delegations_from_ids,
(select array_agg(c.id ORDER BY c.id) from chat_message_t c where c.meeting_user_id = m.id) as chat_message_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_group_meeting_user_ids_meeting_user_t n where n.meeting_user_id = m.id) as group_ids,
(select array_agg(n.structure_level_id ORDER BY n.structure_level_id) from nm_meeting_user_structure_level_ids_structure_level_t n where n.meeting_user_id = m.id) as structure_level_ids
FROM meeting_user_t m;


CREATE VIEW "gender" AS SELECT *,
(select array_agg(u.id ORDER BY u.id) from user_t u where u.gender_id = g.id) as user_ids
FROM gender_t g;


CREATE VIEW "organization_tag" AS SELECT *,
(select array_agg(g.tagged_id ORDER BY g.tagged_id) from gm_organization_tag_tagged_ids_t g where g.organization_tag_id = o.id) as tagged_ids
FROM organization_tag_t o;


CREATE VIEW "theme" AS SELECT *,
(select o.id from organization_t o where o.theme_id = t.id) as theme_for_organization_id
FROM theme_t t;


CREATE VIEW "committee" AS SELECT *,
(select array_agg(m.id ORDER BY m.id) from meeting_t m where m.committee_id = c.id) as meeting_ids,
(
  SELECT array_agg(DISTINCT user_id ORDER BY user_id)
  FROM (
    -- Select user_ids from committees meetings
    SELECT mu.user_id
    FROM meeting_t AS m
    INNER JOIN meeting_user_t AS mu ON mu.meeting_id = m.id
    INNER JOIN nm_group_meeting_user_ids_meeting_user_t AS gmu ON mu.id = gmu.meeting_user_id
    WHERE m.committee_id = c.id

    UNION

    -- Select user_ids from committee managers
    SELECT cmu.user_id
    FROM nm_committee_manager_ids_user_t cmu
    WHERE cmu.committee_id = c.id

    UNION

    -- Select user_id from home committees
    SELECT u.id FROM user_t u WHERE u.home_committee_id = c.id
  ) _
) AS user_ids
,
(select array_agg(n.user_id ORDER BY n.user_id) from nm_committee_manager_ids_user_t n where n.committee_id = c.id) as manager_ids,
(select array_agg(ct.id ORDER BY ct.id) from committee_t ct where ct.parent_id = c.id) as child_ids,
(select array_agg(n.all_parent_id ORDER BY n.all_parent_id) from nm_committee_all_child_ids_committee_t n where n.all_child_id = c.id) as all_parent_ids,
(select array_agg(n.all_child_id ORDER BY n.all_child_id) from nm_committee_all_child_ids_committee_t n where n.all_parent_id = c.id) as all_child_ids,
(select array_agg(u.id ORDER BY u.id) from user_t u where u.home_committee_id = c.id) as native_user_ids,
(select array_agg(n.forward_to_committee_id ORDER BY n.forward_to_committee_id) from nm_committee_forward_to_committee_ids_committee_t n where n.receive_forwardings_from_committee_id = c.id) as forward_to_committee_ids,
(select array_agg(n.receive_forwardings_from_committee_id ORDER BY n.receive_forwardings_from_committee_id) from nm_committee_forward_to_committee_ids_committee_t n where n.forward_to_committee_id = c.id) as receive_forwardings_from_committee_ids,
(select array_agg(g.organization_tag_id ORDER BY g.organization_tag_id) from gm_organization_tag_tagged_ids_t g where g.tagged_id_committee_id = c.id) as organization_tag_ids
FROM committee_t c;

comment on column "committee".user_ids is 'Calculated field: All users which are in a group of a meeting, belonging to the committee or beeing manager of the committee';

CREATE VIEW "meeting" AS SELECT *,
(select array_agg(g.id ORDER BY g.id) from group_t g where g.used_as_motion_poll_default_id = m.id) as motion_poll_default_group_ids,
(select array_agg(p.id ORDER BY p.id) from poll_candidate_list_t p where p.meeting_id = m.id) as poll_candidate_list_ids,
(select array_agg(p.id ORDER BY p.id) from poll_candidate_t p where p.meeting_id = m.id) as poll_candidate_ids,
(select array_agg(mu.id ORDER BY mu.id) from meeting_user_t mu where mu.meeting_id = m.id) as meeting_user_ids,
(select array_agg(g.id ORDER BY g.id) from group_t g where g.used_as_assignment_poll_default_id = m.id) as assignment_poll_default_group_ids,
(select array_agg(g.id ORDER BY g.id) from group_t g where g.used_as_poll_default_id = m.id) as poll_default_group_ids,
(select array_agg(g.id ORDER BY g.id) from group_t g where g.used_as_topic_poll_default_id = m.id) as topic_poll_default_group_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.meeting_id = m.id) as projector_ids,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.meeting_id = m.id) as all_projection_ids,
(select array_agg(p.id ORDER BY p.id) from projector_message_t p where p.meeting_id = m.id) as projector_message_ids,
(select array_agg(p.id ORDER BY p.id) from projector_countdown_t p where p.meeting_id = m.id) as projector_countdown_ids,
(select array_agg(t.id ORDER BY t.id) from tag_t t where t.meeting_id = m.id) as tag_ids,
(select array_agg(a.id ORDER BY a.id) from agenda_item_t a where a.meeting_id = m.id) as agenda_item_ids,
(select array_agg(l.id ORDER BY l.id) from list_of_speakers_t l where l.meeting_id = m.id) as list_of_speakers_ids,
(select array_agg(s.id ORDER BY s.id) from structure_level_list_of_speakers_t s where s.meeting_id = m.id) as structure_level_list_of_speakers_ids,
(select array_agg(p.id ORDER BY p.id) from point_of_order_category_t p where p.meeting_id = m.id) as point_of_order_category_ids,
(select array_agg(s.id ORDER BY s.id) from speaker_t s where s.meeting_id = m.id) as speaker_ids,
(select array_agg(t.id ORDER BY t.id) from topic_t t where t.meeting_id = m.id) as topic_ids,
(select array_agg(g.id ORDER BY g.id) from group_t g where g.meeting_id = m.id) as group_ids,
(select array_agg(mm.id ORDER BY mm.id) from meeting_mediafile_t mm where mm.meeting_id = m.id) as meeting_mediafile_ids,
(select array_agg(mt.id ORDER BY mt.id) from mediafile_t mt where mt.owner_id_meeting_id = m.id) as mediafile_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.meeting_id = m.id) as motion_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.origin_meeting_id = m.id) as forwarded_motion_ids,
(select array_agg(mc.id ORDER BY mc.id) from motion_comment_section_t mc where mc.meeting_id = m.id) as motion_comment_section_ids,
(select array_agg(mc.id ORDER BY mc.id) from motion_category_t mc where mc.meeting_id = m.id) as motion_category_ids,
(select array_agg(mb.id ORDER BY mb.id) from motion_block_t mb where mb.meeting_id = m.id) as motion_block_ids,
(select array_agg(mw.id ORDER BY mw.id) from motion_workflow_t mw where mw.meeting_id = m.id) as motion_workflow_ids,
(select array_agg(mc.id ORDER BY mc.id) from motion_comment_t mc where mc.meeting_id = m.id) as motion_comment_ids,
(select array_agg(ms.id ORDER BY ms.id) from motion_submitter_t ms where ms.meeting_id = m.id) as motion_submitter_ids,
(select array_agg(me.id ORDER BY me.id) from motion_editor_t me where me.meeting_id = m.id) as motion_editor_ids,
(select array_agg(mw.id ORDER BY mw.id) from motion_working_group_speaker_t mw where mw.meeting_id = m.id) as motion_working_group_speaker_ids,
(select array_agg(mc.id ORDER BY mc.id) from motion_change_recommendation_t mc where mc.meeting_id = m.id) as motion_change_recommendation_ids,
(select array_agg(ms.id ORDER BY ms.id) from motion_state_t ms where ms.meeting_id = m.id) as motion_state_ids,
(select array_agg(p.id ORDER BY p.id) from poll_t p where p.meeting_id = m.id) as poll_ids,
(select array_agg(v.id ORDER BY v.id) from vote_t v where v.meeting_id = m.id) as vote_ids,
(select array_agg(a.id ORDER BY a.id) from assignment_t a where a.meeting_id = m.id) as assignment_ids,
(select array_agg(a.id ORDER BY a.id) from assignment_candidate_t a where a.meeting_id = m.id) as assignment_candidate_ids,
(select array_agg(p.id ORDER BY p.id) from personal_note_t p where p.meeting_id = m.id) as personal_note_ids,
(select array_agg(c.id ORDER BY c.id) from chat_group_t c where c.meeting_id = m.id) as chat_group_ids,
(select array_agg(c.id ORDER BY c.id) from chat_message_t c where c.meeting_id = m.id) as chat_message_ids,
(select array_agg(s.id ORDER BY s.id) from structure_level_t s where s.meeting_id = m.id) as structure_level_ids,
(select c.id from committee_t c where c.default_meeting_id = m.id) as default_meeting_for_committee_id,
(select array_agg(g.organization_tag_id ORDER BY g.organization_tag_id) from gm_organization_tag_tagged_ids_t g where g.tagged_id_meeting_id = m.id) as organization_tag_ids,
(select array_agg(n.user_id ORDER BY n.user_id) from nm_meeting_present_user_ids_user_t n where n.meeting_id = m.id) as present_user_ids,
(
  SELECT array_agg(DISTINCT mu.user_id ORDER BY mu.user_id)
  FROM meeting_user_t mu
  INNER JOIN nm_group_meeting_user_ids_meeting_user_t AS gmu ON mu.id = gmu.meeting_user_id
  WHERE mu.meeting_id = m.id
) AS user_ids
,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_meeting_id = m.id) as projection_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_agenda_item_list_in_meeting_id = m.id) as default_projector_agenda_item_list_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_topic_in_meeting_id = m.id) as default_projector_topic_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_list_of_speakers_in_meeting_id = m.id) as default_projector_list_of_speakers_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_current_los_in_meeting_id = m.id) as default_projector_current_los_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_motion_in_meeting_id = m.id) as default_projector_motion_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_amendment_in_meeting_id = m.id) as default_projector_amendment_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_motion_block_in_meeting_id = m.id) as default_projector_motion_block_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_assignment_in_meeting_id = m.id) as default_projector_assignment_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_mediafile_in_meeting_id = m.id) as default_projector_mediafile_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_message_in_meeting_id = m.id) as default_projector_message_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_countdown_in_meeting_id = m.id) as default_projector_countdown_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_assignment_poll_in_meeting_id = m.id) as default_projector_assignment_poll_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_motion_poll_in_meeting_id = m.id) as default_projector_motion_poll_ids,
(select array_agg(p.id ORDER BY p.id) from projector_t p where p.used_as_default_projector_for_poll_in_meeting_id = m.id) as default_projector_poll_ids
FROM meeting_t m;

comment on column "meeting_t".is_active_in_organization_id is 'Backrelation and boolean flag at once';
comment on column "meeting_t".is_archived_in_organization_id is 'Backrelation and boolean flag at once';
comment on column "meeting".user_ids is 'Calculated. All user ids from all users assigned to groups of this meeting.';

CREATE VIEW "structure_level" AS SELECT *,
(select array_agg(n.meeting_user_id ORDER BY n.meeting_user_id) from nm_meeting_user_structure_level_ids_structure_level_t n where n.structure_level_id = s.id) as meeting_user_ids,
(select array_agg(sl.id ORDER BY sl.id) from structure_level_list_of_speakers_t sl where sl.structure_level_id = s.id) as structure_level_list_of_speakers_ids
FROM structure_level_t s;


CREATE VIEW "group" AS SELECT *,
(select array_agg(n.meeting_user_id ORDER BY n.meeting_user_id) from nm_group_meeting_user_ids_meeting_user_t n where n.group_id = g.id) as meeting_user_ids,
(select m.id from meeting_t m where m.default_group_id = g.id) as default_group_for_meeting_id,
(select m.id from meeting_t m where m.admin_group_id = g.id) as admin_group_for_meeting_id,
(select m.id from meeting_t m where m.anonymous_group_id = g.id) as anonymous_group_for_meeting_id,
(select array_agg(n.meeting_mediafile_id ORDER BY n.meeting_mediafile_id) from nm_group_mmagi_meeting_mediafile_t n where n.group_id = g.id) as meeting_mediafile_access_group_ids,
(select array_agg(n.meeting_mediafile_id ORDER BY n.meeting_mediafile_id) from nm_group_mmiagi_meeting_mediafile_t n where n.group_id = g.id) as meeting_mediafile_inherited_access_group_ids,
(select array_agg(n.motion_comment_section_id ORDER BY n.motion_comment_section_id) from nm_group_read_comment_section_ids_motion_comment_section_t n where n.group_id = g.id) as read_comment_section_ids,
(select array_agg(n.motion_comment_section_id ORDER BY n.motion_comment_section_id) from nm_group_write_comment_section_ids_motion_comment_section_t n where n.group_id = g.id) as write_comment_section_ids,
(select array_agg(n.chat_group_id ORDER BY n.chat_group_id) from nm_chat_group_read_group_ids_group_t n where n.group_id = g.id) as read_chat_group_ids,
(select array_agg(n.chat_group_id ORDER BY n.chat_group_id) from nm_chat_group_write_group_ids_group_t n where n.group_id = g.id) as write_chat_group_ids,
(select array_agg(n.poll_id ORDER BY n.poll_id) from nm_group_poll_ids_poll_t n where n.group_id = g.id) as poll_ids
FROM group_t g;

comment on column "group".meeting_mediafile_inherited_access_group_ids is 'Calculated field.';

CREATE VIEW "personal_note" AS SELECT * FROM personal_note_t p;


CREATE VIEW "tag" AS SELECT *,
(select array_agg(g.tagged_id ORDER BY g.tagged_id) from gm_tag_tagged_ids_t g where g.tag_id = t.id) as tagged_ids
FROM tag_t t;


CREATE VIEW "agenda_item" AS SELECT *,
(select array_agg(ai.id ORDER BY ai.id) from agenda_item_t ai where ai.parent_id = a.id) as child_ids,
(select array_agg(g.tag_id ORDER BY g.tag_id) from gm_tag_tagged_ids_t g where g.tagged_id_agenda_item_id = a.id) as tag_ids,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_agenda_item_id = a.id) as projection_ids
FROM agenda_item_t a;


CREATE VIEW "list_of_speakers" AS SELECT *,
(select array_agg(s.id ORDER BY s.id) from speaker_t s where s.list_of_speakers_id = l.id) as speaker_ids,
(select array_agg(s.id ORDER BY s.id) from structure_level_list_of_speakers_t s where s.list_of_speakers_id = l.id) as structure_level_list_of_speakers_ids,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_list_of_speakers_id = l.id) as projection_ids
FROM list_of_speakers_t l;


CREATE VIEW "structure_level_list_of_speakers" AS SELECT *,
(select array_agg(st.id ORDER BY st.id) from speaker_t st where st.structure_level_list_of_speakers_id = s.id) as speaker_ids
FROM structure_level_list_of_speakers_t s;


CREATE VIEW "point_of_order_category" AS SELECT *,
(select array_agg(s.id ORDER BY s.id) from speaker_t s where s.point_of_order_category_id = p.id) as speaker_ids
FROM point_of_order_category_t p;


CREATE VIEW "speaker" AS SELECT * FROM speaker_t s;


CREATE VIEW "topic" AS SELECT *,
(select array_agg(g.meeting_mediafile_id ORDER BY g.meeting_mediafile_id) from gm_meeting_mediafile_attachment_ids_t g where g.attachment_id_topic_id = t.id) as attachment_meeting_mediafile_ids,
(select a.id from agenda_item_t a where a.content_object_id_topic_id = t.id) as agenda_item_id,
(select l.id from list_of_speakers_t l where l.content_object_id_topic_id = t.id) as list_of_speakers_id,
(select array_agg(p.id ORDER BY p.id) from poll_t p where p.content_object_id_topic_id = t.id) as poll_ids,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_topic_id = t.id) as projection_ids
FROM topic_t t;


CREATE VIEW "motion" AS SELECT *,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.lead_motion_id = m.id) as amendment_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.sort_parent_id = m.id) as sort_child_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.origin_id = m.id) as derived_motion_ids,
(select array_agg(n.all_origin_id ORDER BY n.all_origin_id) from nm_motion_all_derived_motion_ids_motion_t n where n.all_derived_motion_id = m.id) as all_origin_ids,
(select array_agg(n.all_derived_motion_id ORDER BY n.all_derived_motion_id) from nm_motion_all_derived_motion_ids_motion_t n where n.all_origin_id = m.id) as all_derived_motion_ids,
(select array_cat((select array_agg(n.identical_motion_id_1 ORDER BY n.identical_motion_id_1) from nm_motion_identical_motion_ids_motion_t n where n.identical_motion_id_2 = m.id), (select array_agg(n.identical_motion_id_2 ORDER BY n.identical_motion_id_2) from nm_motion_identical_motion_ids_motion_t n where n.identical_motion_id_1 = m.id))) as identical_motion_ids,
(select array_agg(g.state_extension_reference_id ORDER BY g.state_extension_reference_id) from gm_motion_state_extension_reference_ids_t g where g.motion_id = m.id) as state_extension_reference_ids,
(select array_agg(g.motion_id ORDER BY g.motion_id) from gm_motion_state_extension_reference_ids_t g where g.state_extension_reference_id_motion_id = m.id) as referenced_in_motion_state_extension_ids,
(select array_agg(g.recommendation_extension_reference_id ORDER BY g.recommendation_extension_reference_id) from gm_motion_recommendation_extension_reference_ids_t g where g.motion_id = m.id) as recommendation_extension_reference_ids,
(select array_agg(g.motion_id ORDER BY g.motion_id) from gm_motion_recommendation_extension_reference_ids_t g where g.recommendation_extension_reference_id_motion_id = m.id) as referenced_in_motion_recommendation_extension_ids,
(select array_agg(ms.id ORDER BY ms.id) from motion_submitter_t ms where ms.motion_id = m.id) as submitter_ids,
(select array_agg(n.meeting_user_id ORDER BY n.meeting_user_id) from nm_meeting_user_supported_motion_ids_motion_t n where n.motion_id = m.id) as supporter_meeting_user_ids,
(select array_agg(me.id ORDER BY me.id) from motion_editor_t me where me.motion_id = m.id) as editor_ids,
(select array_agg(mw.id ORDER BY mw.id) from motion_working_group_speaker_t mw where mw.motion_id = m.id) as working_group_speaker_ids,
(select array_agg(p.id ORDER BY p.id) from poll_t p where p.content_object_id_motion_id = m.id) as poll_ids,
(select array_agg(mc.id ORDER BY mc.id) from motion_change_recommendation_t mc where mc.motion_id = m.id) as change_recommendation_ids,
(select array_agg(mc.id ORDER BY mc.id) from motion_comment_t mc where mc.motion_id = m.id) as comment_ids,
(select a.id from agenda_item_t a where a.content_object_id_motion_id = m.id) as agenda_item_id,
(select l.id from list_of_speakers_t l where l.content_object_id_motion_id = m.id) as list_of_speakers_id,
(select array_agg(g.tag_id ORDER BY g.tag_id) from gm_tag_tagged_ids_t g where g.tagged_id_motion_id = m.id) as tag_ids,
(select array_agg(g.meeting_mediafile_id ORDER BY g.meeting_mediafile_id) from gm_meeting_mediafile_attachment_ids_t g where g.attachment_id_motion_id = m.id) as attachment_meeting_mediafile_ids,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_motion_id = m.id) as projection_ids,
(select array_agg(p.id ORDER BY p.id) from personal_note_t p where p.content_object_id_motion_id = m.id) as personal_note_ids
FROM motion_t m;


CREATE VIEW "motion_submitter" AS SELECT * FROM motion_submitter_t m;


CREATE VIEW "motion_editor" AS SELECT * FROM motion_editor_t m;


CREATE VIEW "motion_working_group_speaker" AS SELECT * FROM motion_working_group_speaker_t m;


CREATE VIEW "motion_comment" AS SELECT * FROM motion_comment_t m;


CREATE VIEW "motion_comment_section" AS SELECT *,
(select array_agg(mc.id ORDER BY mc.id) from motion_comment_t mc where mc.section_id = m.id) as comment_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_group_read_comment_section_ids_motion_comment_section_t n where n.motion_comment_section_id = m.id) as read_group_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_group_write_comment_section_ids_motion_comment_section_t n where n.motion_comment_section_id = m.id) as write_group_ids
FROM motion_comment_section_t m;


CREATE VIEW "motion_category" AS SELECT *,
(select array_agg(mc.id ORDER BY mc.id) from motion_category_t mc where mc.parent_id = m.id) as child_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.category_id = m.id) as motion_ids
FROM motion_category_t m;


CREATE VIEW "motion_block" AS SELECT *,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.block_id = m.id) as motion_ids,
(select a.id from agenda_item_t a where a.content_object_id_motion_block_id = m.id) as agenda_item_id,
(select l.id from list_of_speakers_t l where l.content_object_id_motion_block_id = m.id) as list_of_speakers_id,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_motion_block_id = m.id) as projection_ids
FROM motion_block_t m;


CREATE VIEW "motion_change_recommendation" AS SELECT * FROM motion_change_recommendation_t m;


CREATE VIEW "motion_state" AS SELECT *,
(select array_agg(ms.id ORDER BY ms.id) from motion_state_t ms where ms.submitter_withdraw_state_id = m.id) as submitter_withdraw_back_ids,
(select array_agg(n.next_state_id ORDER BY n.next_state_id) from nm_motion_state_next_state_ids_motion_state_t n where n.previous_state_id = m.id) as next_state_ids,
(select array_agg(n.previous_state_id ORDER BY n.previous_state_id) from nm_motion_state_next_state_ids_motion_state_t n where n.next_state_id = m.id) as previous_state_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.state_id = m.id) as motion_ids,
(select array_agg(mt.id ORDER BY mt.id) from motion_t mt where mt.recommendation_id = m.id) as motion_recommendation_ids,
(select mw.id from motion_workflow_t mw where mw.first_state_id = m.id) as first_state_of_workflow_id
FROM motion_state_t m;


CREATE VIEW "motion_workflow" AS SELECT *,
(select array_agg(ms.id ORDER BY ms.id) from motion_state_t ms where ms.workflow_id = m.id) as state_ids,
(select m1.id from meeting_t m1 where m1.motions_default_workflow_id = m.id) as default_workflow_meeting_id,
(select m1.id from meeting_t m1 where m1.motions_default_amendment_workflow_id = m.id) as default_amendment_workflow_meeting_id
FROM motion_workflow_t m;


CREATE VIEW "poll" AS SELECT *,
(select array_agg(v.id ORDER BY v.id) from vote_t v where v.poll_id = p.id) as vote_ids,
(select array_agg(n.user_id ORDER BY n.user_id) from nm_poll_voted_ids_user_t n where n.poll_id = p.id) as voted_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_group_poll_ids_poll_t n where n.poll_id = p.id) as entitled_group_ids,
(select array_agg(pt.id ORDER BY pt.id) from projection_t pt where pt.content_object_id_poll_id = p.id) as projection_ids
FROM poll_t p;


CREATE VIEW "vote" AS SELECT * FROM vote_t v;


CREATE VIEW "assignment" AS SELECT *,
(select array_agg(ac.id ORDER BY ac.id) from assignment_candidate_t ac where ac.assignment_id = a.id) as candidate_ids,
(select array_agg(p.id ORDER BY p.id) from poll_t p where p.content_object_id_assignment_id = a.id) as poll_ids,
(select ai.id from agenda_item_t ai where ai.content_object_id_assignment_id = a.id) as agenda_item_id,
(select l.id from list_of_speakers_t l where l.content_object_id_assignment_id = a.id) as list_of_speakers_id,
(select array_agg(g.tag_id ORDER BY g.tag_id) from gm_tag_tagged_ids_t g where g.tagged_id_assignment_id = a.id) as tag_ids,
(select array_agg(g.meeting_mediafile_id ORDER BY g.meeting_mediafile_id) from gm_meeting_mediafile_attachment_ids_t g where g.attachment_id_assignment_id = a.id) as attachment_meeting_mediafile_ids,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_assignment_id = a.id) as projection_ids
FROM assignment_t a;


CREATE VIEW "assignment_candidate" AS SELECT * FROM assignment_candidate_t a;


CREATE VIEW "poll_candidate_list" AS SELECT *,
(select array_agg(pc.id ORDER BY pc.id) from poll_candidate_t pc where pc.poll_candidate_list_id = p.id) as poll_candidate_ids
FROM poll_candidate_list_t p;


CREATE VIEW "poll_candidate" AS SELECT * FROM poll_candidate_t p;


CREATE VIEW "mediafile" AS SELECT *,
(select array_agg(mt.id ORDER BY mt.id) from mediafile_t mt where mt.parent_id = m.id) as child_ids,
(select array_agg(mm.id ORDER BY mm.id) from meeting_mediafile_t mm where mm.mediafile_id = m.id) as meeting_mediafile_ids
FROM mediafile_t m;


CREATE VIEW "meeting_mediafile" AS SELECT *,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_group_mmiagi_meeting_mediafile_t n where n.meeting_mediafile_id = m.id) as inherited_access_group_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_group_mmagi_meeting_mediafile_t n where n.meeting_mediafile_id = m.id) as access_group_ids,
(select l.id from list_of_speakers_t l where l.content_object_id_meeting_mediafile_id = m.id) as list_of_speakers_id,
(select array_agg(p.id ORDER BY p.id) from projection_t p where p.content_object_id_meeting_mediafile_id = m.id) as projection_ids,
(select array_agg(g.attachment_id ORDER BY g.attachment_id) from gm_meeting_mediafile_attachment_ids_t g where g.meeting_mediafile_id = m.id) as attachment_ids,
(select m1.id from meeting_t m1 where m1.logo_projector_main_id = m.id) as used_as_logo_projector_main_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_projector_header_id = m.id) as used_as_logo_projector_header_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_web_header_id = m.id) as used_as_logo_web_header_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_pdf_header_l_id = m.id) as used_as_logo_pdf_header_l_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_pdf_header_r_id = m.id) as used_as_logo_pdf_header_r_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_pdf_footer_l_id = m.id) as used_as_logo_pdf_footer_l_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_pdf_footer_r_id = m.id) as used_as_logo_pdf_footer_r_in_meeting_id,
(select m1.id from meeting_t m1 where m1.logo_pdf_ballot_paper_id = m.id) as used_as_logo_pdf_ballot_paper_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_regular_id = m.id) as used_as_font_regular_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_italic_id = m.id) as used_as_font_italic_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_bold_id = m.id) as used_as_font_bold_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_bold_italic_id = m.id) as used_as_font_bold_italic_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_monospace_id = m.id) as used_as_font_monospace_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_chyron_speaker_name_id = m.id) as used_as_font_chyron_speaker_name_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_projector_h1_id = m.id) as used_as_font_projector_h1_in_meeting_id,
(select m1.id from meeting_t m1 where m1.font_projector_h2_id = m.id) as used_as_font_projector_h2_in_meeting_id
FROM meeting_mediafile_t m;

comment on column "meeting_mediafile".inherited_access_group_ids is 'Calculated in actions. Shows what access group permissions are actually relevant. Calculated as the intersection of this meeting_mediafiles access_group_ids and the related mediafiles potential parent mediafiles inherited_access_group_ids. If the parent has no meeting_mediafile for this meeting, its inherited access group is assumed to be the meetings admin group. If there is no parent, the inherited_access_group_ids is equal to the access_group_ids. If the access_group_ids are empty, the interpretations is that every group has access rights, therefore the parent inherited_access_group_ids are used as-is.';

CREATE VIEW "projector" AS SELECT *,
(select array_agg(pt.id ORDER BY pt.id) from projection_t pt where pt.current_projector_id = p.id) as current_projection_ids,
(select array_agg(pt.id ORDER BY pt.id) from projection_t pt where pt.preview_projector_id = p.id) as preview_projection_ids,
(select array_agg(pt.id ORDER BY pt.id) from projection_t pt where pt.history_projector_id = p.id) as history_projection_ids,
(select m.id from meeting_t m where m.reference_projector_id = p.id) as used_as_reference_projector_meeting_id
FROM projector_t p;


CREATE VIEW "projection" AS SELECT * FROM projection_t p;


CREATE VIEW "projector_message" AS SELECT *,
(select array_agg(pt.id ORDER BY pt.id) from projection_t pt where pt.content_object_id_projector_message_id = p.id) as projection_ids
FROM projector_message_t p;


CREATE VIEW "projector_countdown" AS SELECT *,
(select array_agg(pt.id ORDER BY pt.id) from projection_t pt where pt.content_object_id_projector_countdown_id = p.id) as projection_ids,
(select m.id from meeting_t m where m.list_of_speakers_countdown_id = p.id) as used_as_list_of_speakers_countdown_meeting_id,
(select m.id from meeting_t m where m.poll_countdown_id = p.id) as used_as_poll_countdown_meeting_id
FROM projector_countdown_t p;


CREATE VIEW "chat_group" AS SELECT *,
(select array_agg(cm.id ORDER BY cm.id) from chat_message_t cm where cm.chat_group_id = c.id) as chat_message_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_chat_group_read_group_ids_group_t n where n.chat_group_id = c.id) as read_group_ids,
(select array_agg(n.group_id ORDER BY n.group_id) from nm_chat_group_write_group_ids_group_t n where n.chat_group_id = c.id) as write_group_ids
FROM chat_group_t c;


CREATE VIEW "chat_message" AS SELECT * FROM chat_message_t c;


CREATE VIEW "action_worker" AS SELECT * FROM action_worker_t a;


CREATE VIEW "import_preview" AS SELECT * FROM import_preview_t i;



-- Alter table relations
ALTER TABLE organization_t ADD FOREIGN KEY(theme_id) REFERENCES theme_t(id) INITIALLY DEFERRED;

ALTER TABLE user_t ADD FOREIGN KEY(gender_id) REFERENCES gender_t(id) INITIALLY DEFERRED;
ALTER TABLE user_t ADD FOREIGN KEY(home_committee_id) REFERENCES committee_t(id) INITIALLY DEFERRED;

ALTER TABLE meeting_user_t ADD FOREIGN KEY(user_id) REFERENCES user_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_user_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_user_t ADD FOREIGN KEY(vote_delegated_to_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;

ALTER TABLE committee_t ADD FOREIGN KEY(default_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE committee_t ADD FOREIGN KEY(parent_id) REFERENCES committee_t(id) INITIALLY DEFERRED;

ALTER TABLE meeting_t ADD FOREIGN KEY(is_active_in_organization_id) REFERENCES organization_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(is_archived_in_organization_id) REFERENCES organization_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(template_for_organization_id) REFERENCES organization_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(motions_default_workflow_id) REFERENCES motion_workflow_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(motions_default_amendment_workflow_id) REFERENCES motion_workflow_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_projector_main_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_projector_header_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_web_header_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_pdf_header_l_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_pdf_header_r_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_pdf_footer_l_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_pdf_footer_r_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(logo_pdf_ballot_paper_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_regular_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_italic_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_bold_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_bold_italic_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_monospace_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_chyron_speaker_name_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_projector_h1_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(font_projector_h2_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(committee_id) REFERENCES committee_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(reference_projector_id) REFERENCES projector_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(list_of_speakers_countdown_id) REFERENCES projector_countdown_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(poll_countdown_id) REFERENCES projector_countdown_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(default_group_id) REFERENCES group_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(admin_group_id) REFERENCES group_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_t ADD FOREIGN KEY(anonymous_group_id) REFERENCES group_t(id) INITIALLY DEFERRED;

ALTER TABLE structure_level_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE group_t ADD FOREIGN KEY(used_as_motion_poll_default_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE group_t ADD FOREIGN KEY(used_as_assignment_poll_default_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE group_t ADD FOREIGN KEY(used_as_topic_poll_default_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE group_t ADD FOREIGN KEY(used_as_poll_default_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE group_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE personal_note_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE personal_note_t ADD FOREIGN KEY(content_object_id_motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE personal_note_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE tag_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE agenda_item_t ADD FOREIGN KEY(content_object_id_motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE agenda_item_t ADD FOREIGN KEY(content_object_id_motion_block_id) REFERENCES motion_block_t(id) INITIALLY DEFERRED;
ALTER TABLE agenda_item_t ADD FOREIGN KEY(content_object_id_assignment_id) REFERENCES assignment_t(id) INITIALLY DEFERRED;
ALTER TABLE agenda_item_t ADD FOREIGN KEY(content_object_id_topic_id) REFERENCES topic_t(id) INITIALLY DEFERRED;
ALTER TABLE agenda_item_t ADD FOREIGN KEY(parent_id) REFERENCES agenda_item_t(id) INITIALLY DEFERRED;
ALTER TABLE agenda_item_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE list_of_speakers_t ADD FOREIGN KEY(content_object_id_motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE list_of_speakers_t ADD FOREIGN KEY(content_object_id_motion_block_id) REFERENCES motion_block_t(id) INITIALLY DEFERRED;
ALTER TABLE list_of_speakers_t ADD FOREIGN KEY(content_object_id_assignment_id) REFERENCES assignment_t(id) INITIALLY DEFERRED;
ALTER TABLE list_of_speakers_t ADD FOREIGN KEY(content_object_id_topic_id) REFERENCES topic_t(id) INITIALLY DEFERRED;
ALTER TABLE list_of_speakers_t ADD FOREIGN KEY(content_object_id_meeting_mediafile_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE list_of_speakers_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE structure_level_list_of_speakers_t ADD FOREIGN KEY(structure_level_id) REFERENCES structure_level_t(id) INITIALLY DEFERRED;
ALTER TABLE structure_level_list_of_speakers_t ADD FOREIGN KEY(list_of_speakers_id) REFERENCES list_of_speakers_t(id) INITIALLY DEFERRED;
ALTER TABLE structure_level_list_of_speakers_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE point_of_order_category_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE speaker_t ADD FOREIGN KEY(list_of_speakers_id) REFERENCES list_of_speakers_t(id) INITIALLY DEFERRED;
ALTER TABLE speaker_t ADD FOREIGN KEY(structure_level_list_of_speakers_id) REFERENCES structure_level_list_of_speakers_t(id) INITIALLY DEFERRED;
ALTER TABLE speaker_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE speaker_t ADD FOREIGN KEY(point_of_order_category_id) REFERENCES point_of_order_category_t(id) INITIALLY DEFERRED;
ALTER TABLE speaker_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE topic_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_t ADD FOREIGN KEY(lead_motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(sort_parent_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(origin_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(origin_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(state_id) REFERENCES motion_state_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(recommendation_id) REFERENCES motion_state_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(category_id) REFERENCES motion_category_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(block_id) REFERENCES motion_block_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_submitter_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_submitter_t ADD FOREIGN KEY(motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_submitter_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_editor_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_editor_t ADD FOREIGN KEY(motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_editor_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_working_group_speaker_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_working_group_speaker_t ADD FOREIGN KEY(motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_working_group_speaker_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_comment_t ADD FOREIGN KEY(motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_comment_t ADD FOREIGN KEY(section_id) REFERENCES motion_comment_section_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_comment_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_comment_section_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_category_t ADD FOREIGN KEY(parent_id) REFERENCES motion_category_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_category_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_block_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_change_recommendation_t ADD FOREIGN KEY(motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_change_recommendation_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_state_t ADD FOREIGN KEY(submitter_withdraw_state_id) REFERENCES motion_state_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_state_t ADD FOREIGN KEY(workflow_id) REFERENCES motion_workflow_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_state_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE motion_workflow_t ADD FOREIGN KEY(first_state_id) REFERENCES motion_state_t(id) INITIALLY DEFERRED;
ALTER TABLE motion_workflow_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE poll_t ADD FOREIGN KEY(content_object_id_motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE poll_t ADD FOREIGN KEY(content_object_id_assignment_id) REFERENCES assignment_t(id) INITIALLY DEFERRED;
ALTER TABLE poll_t ADD FOREIGN KEY(content_object_id_topic_id) REFERENCES topic_t(id) INITIALLY DEFERRED;
ALTER TABLE poll_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE vote_t ADD FOREIGN KEY(poll_id) REFERENCES poll_t(id) INITIALLY DEFERRED;
ALTER TABLE vote_t ADD FOREIGN KEY(acting_user_id) REFERENCES user_t(id) INITIALLY DEFERRED;
ALTER TABLE vote_t ADD FOREIGN KEY(represented_user_id) REFERENCES user_t(id) INITIALLY DEFERRED;
ALTER TABLE vote_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE assignment_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE assignment_candidate_t ADD FOREIGN KEY(assignment_id) REFERENCES assignment_t(id) INITIALLY DEFERRED;
ALTER TABLE assignment_candidate_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE assignment_candidate_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE poll_candidate_list_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE poll_candidate_t ADD FOREIGN KEY(poll_candidate_list_id) REFERENCES poll_candidate_list_t(id) INITIALLY DEFERRED;
ALTER TABLE poll_candidate_t ADD FOREIGN KEY(user_id) REFERENCES user_t(id) INITIALLY DEFERRED;
ALTER TABLE poll_candidate_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE mediafile_t ADD FOREIGN KEY(published_to_meetings_in_organization_id) REFERENCES organization_t(id) INITIALLY DEFERRED;
ALTER TABLE mediafile_t ADD FOREIGN KEY(parent_id) REFERENCES mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE mediafile_t ADD FOREIGN KEY(owner_id_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE mediafile_t ADD FOREIGN KEY(owner_id_organization_id) REFERENCES organization_t(id) INITIALLY DEFERRED;

ALTER TABLE meeting_mediafile_t ADD FOREIGN KEY(mediafile_id) REFERENCES mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE meeting_mediafile_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_agenda_item_list_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_topic_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_list_of_speakers_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_current_los_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_motion_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_amendment_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_motion_block_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_assignment_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_mediafile_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_message_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_countdown_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_assignment_poll_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_motion_poll_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(used_as_default_projector_for_poll_in_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projector_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE projection_t ADD FOREIGN KEY(current_projector_id) REFERENCES projector_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(preview_projector_id) REFERENCES projector_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(history_projector_id) REFERENCES projector_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_motion_id) REFERENCES motion_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_meeting_mediafile_id) REFERENCES meeting_mediafile_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_list_of_speakers_id) REFERENCES list_of_speakers_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_motion_block_id) REFERENCES motion_block_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_assignment_id) REFERENCES assignment_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_agenda_item_id) REFERENCES agenda_item_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_topic_id) REFERENCES topic_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_poll_id) REFERENCES poll_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_projector_message_id) REFERENCES projector_message_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(content_object_id_projector_countdown_id) REFERENCES projector_countdown_t(id) INITIALLY DEFERRED;
ALTER TABLE projection_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE projector_message_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE projector_countdown_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE chat_group_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;

ALTER TABLE chat_message_t ADD FOREIGN KEY(meeting_user_id) REFERENCES meeting_user_t(id) INITIALLY DEFERRED;
ALTER TABLE chat_message_t ADD FOREIGN KEY(chat_group_id) REFERENCES chat_group_t(id) INITIALLY DEFERRED;
ALTER TABLE chat_message_t ADD FOREIGN KEY(meeting_id) REFERENCES meeting_t(id) INITIALLY DEFERRED;



-- Create triggers checking foreign_id not null for relation-lists

-- definition trigger not null for meeting.default_projector_agenda_item_list_ids against projector_t.used_as_default_projector_for_agenda_item_list_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_agenda_item_list_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_agenda_item_list_ids', 'used_as_default_projector_for_agenda_item_list_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_agenda_item_list_ids AFTER UPDATE OF used_as_default_projector_for_agenda_item_list_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_agenda_item_list_ids', 'used_as_default_projector_for_agenda_item_list_in_meeting_id');


-- definition trigger not null for meeting.default_projector_topic_ids against projector_t.used_as_default_projector_for_topic_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_topic_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_topic_ids', 'used_as_default_projector_for_topic_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_topic_ids AFTER UPDATE OF used_as_default_projector_for_topic_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_topic_ids', 'used_as_default_projector_for_topic_in_meeting_id');


-- definition trigger not null for meeting.default_projector_list_of_speakers_ids against projector_t.used_as_default_projector_for_list_of_speakers_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_list_of_speakers_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_list_of_speakers_ids', 'used_as_default_projector_for_list_of_speakers_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_list_of_speakers_ids AFTER UPDATE OF used_as_default_projector_for_list_of_speakers_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_list_of_speakers_ids', 'used_as_default_projector_for_list_of_speakers_in_meeting_id');


-- definition trigger not null for meeting.default_projector_current_los_ids against projector_t.used_as_default_projector_for_current_los_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_current_los_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_current_los_ids', 'used_as_default_projector_for_current_los_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_current_los_ids AFTER UPDATE OF used_as_default_projector_for_current_los_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_current_los_ids', 'used_as_default_projector_for_current_los_in_meeting_id');


-- definition trigger not null for meeting.default_projector_motion_ids against projector_t.used_as_default_projector_for_motion_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_motion_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_motion_ids', 'used_as_default_projector_for_motion_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_motion_ids AFTER UPDATE OF used_as_default_projector_for_motion_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_motion_ids', 'used_as_default_projector_for_motion_in_meeting_id');


-- definition trigger not null for meeting.default_projector_amendment_ids against projector_t.used_as_default_projector_for_amendment_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_amendment_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_amendment_ids', 'used_as_default_projector_for_amendment_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_amendment_ids AFTER UPDATE OF used_as_default_projector_for_amendment_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_amendment_ids', 'used_as_default_projector_for_amendment_in_meeting_id');


-- definition trigger not null for meeting.default_projector_motion_block_ids against projector_t.used_as_default_projector_for_motion_block_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_motion_block_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_motion_block_ids', 'used_as_default_projector_for_motion_block_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_motion_block_ids AFTER UPDATE OF used_as_default_projector_for_motion_block_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_motion_block_ids', 'used_as_default_projector_for_motion_block_in_meeting_id');


-- definition trigger not null for meeting.default_projector_assignment_ids against projector_t.used_as_default_projector_for_assignment_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_assignment_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_assignment_ids', 'used_as_default_projector_for_assignment_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_assignment_ids AFTER UPDATE OF used_as_default_projector_for_assignment_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_assignment_ids', 'used_as_default_projector_for_assignment_in_meeting_id');


-- definition trigger not null for meeting.default_projector_mediafile_ids against projector_t.used_as_default_projector_for_mediafile_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_mediafile_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_mediafile_ids', 'used_as_default_projector_for_mediafile_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_mediafile_ids AFTER UPDATE OF used_as_default_projector_for_mediafile_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_mediafile_ids', 'used_as_default_projector_for_mediafile_in_meeting_id');


-- definition trigger not null for meeting.default_projector_message_ids against projector_t.used_as_default_projector_for_message_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_message_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_message_ids', 'used_as_default_projector_for_message_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_message_ids AFTER UPDATE OF used_as_default_projector_for_message_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_message_ids', 'used_as_default_projector_for_message_in_meeting_id');


-- definition trigger not null for meeting.default_projector_countdown_ids against projector_t.used_as_default_projector_for_countdown_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_countdown_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_countdown_ids', 'used_as_default_projector_for_countdown_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_countdown_ids AFTER UPDATE OF used_as_default_projector_for_countdown_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_countdown_ids', 'used_as_default_projector_for_countdown_in_meeting_id');


-- definition trigger not null for meeting.default_projector_assignment_poll_ids against projector_t.used_as_default_projector_for_assignment_poll_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_assignment_poll_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_assignment_poll_ids', 'used_as_default_projector_for_assignment_poll_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_assignment_poll_ids AFTER UPDATE OF used_as_default_projector_for_assignment_poll_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_assignment_poll_ids', 'used_as_default_projector_for_assignment_poll_in_meeting_id');


-- definition trigger not null for meeting.default_projector_motion_poll_ids against projector_t.used_as_default_projector_for_motion_poll_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_motion_poll_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_motion_poll_ids', 'used_as_default_projector_for_motion_poll_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_motion_poll_ids AFTER UPDATE OF used_as_default_projector_for_motion_poll_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_motion_poll_ids', 'used_as_default_projector_for_motion_poll_in_meeting_id');


-- definition trigger not null for meeting.default_projector_poll_ids against projector_t.used_as_default_projector_for_poll_in_meeting_id
CREATE CONSTRAINT TRIGGER tr_i_meeting_default_projector_poll_ids AFTER INSERT ON projector_t INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_poll_ids', 'used_as_default_projector_for_poll_in_meeting_id');

CREATE CONSTRAINT TRIGGER tr_ud_meeting_default_projector_poll_ids AFTER UPDATE OF used_as_default_projector_for_poll_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION check_not_null_for_relation_lists('meeting', 'default_projector_poll_ids', 'used_as_default_projector_for_poll_in_meeting_id');




-- Create triggers preventing mirrored duplicates in fields referencing themselves

-- definition trigger unique ids pair for motion.identical_motion_ids
CREATE TRIGGER restrict_motion_identical_motion_ids BEFORE INSERT OR UPDATE ON nm_motion_identical_motion_ids_motion_t
FOR EACH ROW EXECUTE FUNCTION check_unique_ids_pair('identical_motion_id');




-- Create triggers for notify
CREATE TRIGGER tr_log_organization AFTER INSERT OR UPDATE OR DELETE ON organization_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('organization');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON organization_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_organization_t_theme_id AFTER INSERT OR UPDATE OF theme_id OR DELETE ON organization_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('theme', 'theme_id');

CREATE TRIGGER tr_log_user AFTER INSERT OR UPDATE OR DELETE ON user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('user');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON user_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_user_t_gender_id AFTER INSERT OR UPDATE OF gender_id OR DELETE ON user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('gender', 'gender_id');
CREATE TRIGGER tr_log_user_t_home_committee_id AFTER INSERT OR UPDATE OF home_committee_id OR DELETE ON user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('committee', 'home_committee_id');

CREATE TRIGGER tr_log_meeting_user AFTER INSERT OR UPDATE OR DELETE ON meeting_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('meeting_user');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON meeting_user_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_meeting_user_t_user_id AFTER INSERT OR UPDATE OF user_id OR DELETE ON meeting_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('user', 'user_id');
CREATE TRIGGER tr_log_meeting_user_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON meeting_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_nm_meeting_user_supported_motion_ids_motion_t AFTER INSERT OR UPDATE OR DELETE ON nm_meeting_user_supported_motion_ids_motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user','meeting_user_id','motion','motion_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_meeting_user_supported_motion_ids_motion_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_meeting_user_t_vote_delegated_to_id AFTER INSERT OR UPDATE OF vote_delegated_to_id OR DELETE ON meeting_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'vote_delegated_to_id');

CREATE TRIGGER tr_log_nm_meeting_user_structure_level_ids_structure_level_t AFTER INSERT OR UPDATE OR DELETE ON nm_meeting_user_structure_level_ids_structure_level_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user','meeting_user_id','structure_level','structure_level_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_meeting_user_structure_level_ids_structure_level_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_gender AFTER INSERT OR UPDATE OR DELETE ON gender_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('gender');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON gender_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_organization_tag AFTER INSERT OR UPDATE OR DELETE ON organization_tag_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('organization_tag');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON organization_tag_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_tagged_id_committee_id_gm_organization_tag_tagged_ids_t AFTER INSERT OR UPDATE OF tagged_id_committee_id OR DELETE ON gm_organization_tag_tagged_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization_tag','organization_tag_id','committee','tagged_id_committee_id');

CREATE TRIGGER tr_log_tagged_id_meeting_id_gm_organization_tag_tagged_ids_t AFTER INSERT OR UPDATE OF tagged_id_meeting_id OR DELETE ON gm_organization_tag_tagged_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization_tag','organization_tag_id','meeting','tagged_id_meeting_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON gm_organization_tag_tagged_ids_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_theme AFTER INSERT OR UPDATE OR DELETE ON theme_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('theme');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON theme_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_committee AFTER INSERT OR UPDATE OR DELETE ON committee_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('committee');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON committee_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_committee_t_default_meeting_id AFTER INSERT OR UPDATE OF default_meeting_id OR DELETE ON committee_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'default_meeting_id');

CREATE TRIGGER tr_log_nm_committee_manager_ids_user_t AFTER INSERT OR UPDATE OR DELETE ON nm_committee_manager_ids_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('committee','committee_id','user','user_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_committee_manager_ids_user_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_committee_t_parent_id AFTER INSERT OR UPDATE OF parent_id OR DELETE ON committee_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('committee', 'parent_id');

CREATE TRIGGER tr_log_nm_committee_all_child_ids_committee_t AFTER INSERT OR UPDATE OR DELETE ON nm_committee_all_child_ids_committee_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('committee','all_child_id','committee','all_parent_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_committee_all_child_ids_committee_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_committee_forward_to_committee_ids_committee_t AFTER INSERT OR UPDATE OR DELETE ON nm_committee_forward_to_committee_ids_committee_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('committee','forward_to_committee_id','committee','receive_forwardings_from_committee_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_committee_forward_to_committee_ids_committee_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_meeting AFTER INSERT OR UPDATE OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('meeting');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON meeting_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_meeting_t_is_active_in_organization_id AFTER INSERT OR UPDATE OF is_active_in_organization_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization', 'is_active_in_organization_id');
CREATE TRIGGER tr_log_meeting_t_is_archived_in_organization_id AFTER INSERT OR UPDATE OF is_archived_in_organization_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization', 'is_archived_in_organization_id');
CREATE TRIGGER tr_log_meeting_t_template_for_organization_id AFTER INSERT OR UPDATE OF template_for_organization_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization', 'template_for_organization_id');
CREATE TRIGGER tr_log_meeting_t_motions_default_workflow_id AFTER INSERT OR UPDATE OF motions_default_workflow_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_workflow', 'motions_default_workflow_id');
CREATE TRIGGER tr_log_meeting_t_motions_default_amendment_workflow_id AFTER INSERT OR UPDATE OF motions_default_amendment_workflow_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_workflow', 'motions_default_amendment_workflow_id');
CREATE TRIGGER tr_log_meeting_t_logo_projector_main_id AFTER INSERT OR UPDATE OF logo_projector_main_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_projector_main_id');
CREATE TRIGGER tr_log_meeting_t_logo_projector_header_id AFTER INSERT OR UPDATE OF logo_projector_header_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_projector_header_id');
CREATE TRIGGER tr_log_meeting_t_logo_web_header_id AFTER INSERT OR UPDATE OF logo_web_header_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_web_header_id');
CREATE TRIGGER tr_log_meeting_t_logo_pdf_header_l_id AFTER INSERT OR UPDATE OF logo_pdf_header_l_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_pdf_header_l_id');
CREATE TRIGGER tr_log_meeting_t_logo_pdf_header_r_id AFTER INSERT OR UPDATE OF logo_pdf_header_r_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_pdf_header_r_id');
CREATE TRIGGER tr_log_meeting_t_logo_pdf_footer_l_id AFTER INSERT OR UPDATE OF logo_pdf_footer_l_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_pdf_footer_l_id');
CREATE TRIGGER tr_log_meeting_t_logo_pdf_footer_r_id AFTER INSERT OR UPDATE OF logo_pdf_footer_r_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_pdf_footer_r_id');
CREATE TRIGGER tr_log_meeting_t_logo_pdf_ballot_paper_id AFTER INSERT OR UPDATE OF logo_pdf_ballot_paper_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'logo_pdf_ballot_paper_id');
CREATE TRIGGER tr_log_meeting_t_font_regular_id AFTER INSERT OR UPDATE OF font_regular_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_regular_id');
CREATE TRIGGER tr_log_meeting_t_font_italic_id AFTER INSERT OR UPDATE OF font_italic_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_italic_id');
CREATE TRIGGER tr_log_meeting_t_font_bold_id AFTER INSERT OR UPDATE OF font_bold_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_bold_id');
CREATE TRIGGER tr_log_meeting_t_font_bold_italic_id AFTER INSERT OR UPDATE OF font_bold_italic_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_bold_italic_id');
CREATE TRIGGER tr_log_meeting_t_font_monospace_id AFTER INSERT OR UPDATE OF font_monospace_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_monospace_id');
CREATE TRIGGER tr_log_meeting_t_font_chyron_speaker_name_id AFTER INSERT OR UPDATE OF font_chyron_speaker_name_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_chyron_speaker_name_id');
CREATE TRIGGER tr_log_meeting_t_font_projector_h1_id AFTER INSERT OR UPDATE OF font_projector_h1_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_projector_h1_id');
CREATE TRIGGER tr_log_meeting_t_font_projector_h2_id AFTER INSERT OR UPDATE OF font_projector_h2_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile', 'font_projector_h2_id');
CREATE TRIGGER tr_log_meeting_t_committee_id AFTER INSERT OR UPDATE OF committee_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('committee', 'committee_id');

CREATE TRIGGER tr_log_nm_meeting_present_user_ids_user_t AFTER INSERT OR UPDATE OR DELETE ON nm_meeting_present_user_ids_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting','meeting_id','user','user_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_meeting_present_user_ids_user_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_meeting_t_reference_projector_id AFTER INSERT OR UPDATE OF reference_projector_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector', 'reference_projector_id');
CREATE TRIGGER tr_log_meeting_t_list_of_speakers_countdown_id AFTER INSERT OR UPDATE OF list_of_speakers_countdown_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector_countdown', 'list_of_speakers_countdown_id');
CREATE TRIGGER tr_log_meeting_t_poll_countdown_id AFTER INSERT OR UPDATE OF poll_countdown_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector_countdown', 'poll_countdown_id');
CREATE TRIGGER tr_log_meeting_t_default_group_id AFTER INSERT OR UPDATE OF default_group_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group', 'default_group_id');
CREATE TRIGGER tr_log_meeting_t_admin_group_id AFTER INSERT OR UPDATE OF admin_group_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group', 'admin_group_id');
CREATE TRIGGER tr_log_meeting_t_anonymous_group_id AFTER INSERT OR UPDATE OF anonymous_group_id OR DELETE ON meeting_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group', 'anonymous_group_id');

CREATE TRIGGER tr_log_structure_level AFTER INSERT OR UPDATE OR DELETE ON structure_level_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('structure_level');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON structure_level_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_structure_level_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON structure_level_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_group AFTER INSERT OR UPDATE OR DELETE ON group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('group');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON group_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_nm_group_meeting_user_ids_meeting_user_t AFTER INSERT OR UPDATE OR DELETE ON nm_group_meeting_user_ids_meeting_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group','group_id','meeting_user','meeting_user_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_group_meeting_user_ids_meeting_user_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_group_mmagi_meeting_mediafile_t AFTER INSERT OR UPDATE OR DELETE ON nm_group_mmagi_meeting_mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group','group_id','meeting_mediafile','meeting_mediafile_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_group_mmagi_meeting_mediafile_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_group_mmiagi_meeting_mediafile_t AFTER INSERT OR UPDATE OR DELETE ON nm_group_mmiagi_meeting_mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group','group_id','meeting_mediafile','meeting_mediafile_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_group_mmiagi_meeting_mediafile_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_group_read_comment_section_ids_motion_comment_section_t AFTER INSERT OR UPDATE OR DELETE ON nm_group_read_comment_section_ids_motion_comment_section_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group','group_id','motion_comment_section','motion_comment_section_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_group_read_comment_section_ids_motion_comment_section_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_group_write_comment_section_ids_motion_comment_section_t AFTER INSERT OR UPDATE OR DELETE ON nm_group_write_comment_section_ids_motion_comment_section_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group','group_id','motion_comment_section','motion_comment_section_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_group_write_comment_section_ids_motion_comment_section_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_group_poll_ids_poll_t AFTER INSERT OR UPDATE OR DELETE ON nm_group_poll_ids_poll_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('group','group_id','poll','poll_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_group_poll_ids_poll_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_group_t_used_as_motion_poll_default_id AFTER INSERT OR UPDATE OF used_as_motion_poll_default_id OR DELETE ON group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_motion_poll_default_id');
CREATE TRIGGER tr_log_group_t_used_as_assignment_poll_default_id AFTER INSERT OR UPDATE OF used_as_assignment_poll_default_id OR DELETE ON group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_assignment_poll_default_id');
CREATE TRIGGER tr_log_group_t_used_as_topic_poll_default_id AFTER INSERT OR UPDATE OF used_as_topic_poll_default_id OR DELETE ON group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_topic_poll_default_id');
CREATE TRIGGER tr_log_group_t_used_as_poll_default_id AFTER INSERT OR UPDATE OF used_as_poll_default_id OR DELETE ON group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_poll_default_id');
CREATE TRIGGER tr_log_group_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_personal_note AFTER INSERT OR UPDATE OR DELETE ON personal_note_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('personal_note');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON personal_note_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_personal_note_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON personal_note_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');

CREATE TRIGGER tr_log_motion_content_object_id_motion_id AFTER INSERT OR UPDATE OF content_object_id_motion_id OR DELETE ON personal_note_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','content_object_id_motion_id');
CREATE TRIGGER tr_log_personal_note_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON personal_note_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_tag AFTER INSERT OR UPDATE OR DELETE ON tag_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('tag');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON tag_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_tagged_id_agenda_item_id_gm_tag_tagged_ids_t AFTER INSERT OR UPDATE OF tagged_id_agenda_item_id OR DELETE ON gm_tag_tagged_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('tag','tag_id','agenda_item','tagged_id_agenda_item_id');

CREATE TRIGGER tr_log_tagged_id_assignment_id_gm_tag_tagged_ids_t AFTER INSERT OR UPDATE OF tagged_id_assignment_id OR DELETE ON gm_tag_tagged_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('tag','tag_id','assignment','tagged_id_assignment_id');

CREATE TRIGGER tr_log_tagged_id_motion_id_gm_tag_tagged_ids_t AFTER INSERT OR UPDATE OF tagged_id_motion_id OR DELETE ON gm_tag_tagged_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('tag','tag_id','motion','tagged_id_motion_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON gm_tag_tagged_ids_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_tag_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON tag_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_agenda_item AFTER INSERT OR UPDATE OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('agenda_item');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON agenda_item_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_motion_content_object_id_motion_id AFTER INSERT OR UPDATE OF content_object_id_motion_id OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','content_object_id_motion_id');

CREATE TRIGGER tr_log_motion_block_content_object_id_motion_block_id AFTER INSERT OR UPDATE OF content_object_id_motion_block_id OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_block','content_object_id_motion_block_id');

CREATE TRIGGER tr_log_assignment_content_object_id_assignment_id AFTER INSERT OR UPDATE OF content_object_id_assignment_id OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('assignment','content_object_id_assignment_id');

CREATE TRIGGER tr_log_topic_content_object_id_topic_id AFTER INSERT OR UPDATE OF content_object_id_topic_id OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('topic','content_object_id_topic_id');
CREATE TRIGGER tr_log_agenda_item_t_parent_id AFTER INSERT OR UPDATE OF parent_id OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('agenda_item', 'parent_id');
CREATE TRIGGER tr_log_agenda_item_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON agenda_item_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_list_of_speakers AFTER INSERT OR UPDATE OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('list_of_speakers');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON list_of_speakers_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_motion_content_object_id_motion_id AFTER INSERT OR UPDATE OF content_object_id_motion_id OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','content_object_id_motion_id');

CREATE TRIGGER tr_log_motion_block_content_object_id_motion_block_id AFTER INSERT OR UPDATE OF content_object_id_motion_block_id OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_block','content_object_id_motion_block_id');

CREATE TRIGGER tr_log_assignment_content_object_id_assignment_id AFTER INSERT OR UPDATE OF content_object_id_assignment_id OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('assignment','content_object_id_assignment_id');

CREATE TRIGGER tr_log_topic_content_object_id_topic_id AFTER INSERT OR UPDATE OF content_object_id_topic_id OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('topic','content_object_id_topic_id');

CREATE TRIGGER tr_log_meeting_mediafile_content_object_id_meeting_mediafile_id AFTER INSERT OR UPDATE OF content_object_id_meeting_mediafile_id OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile','content_object_id_meeting_mediafile_id');
CREATE TRIGGER tr_log_list_of_speakers_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_structure_level_list_of_speakers AFTER INSERT OR UPDATE OR DELETE ON structure_level_list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('structure_level_list_of_speakers');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON structure_level_list_of_speakers_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_structure_level_list_of_speakers_t_structure_level_id AFTER INSERT OR UPDATE OF structure_level_id OR DELETE ON structure_level_list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('structure_level', 'structure_level_id');
CREATE TRIGGER tr_log_structure_level_list_of_speakers_t_list_of_speakers_id AFTER INSERT OR UPDATE OF list_of_speakers_id OR DELETE ON structure_level_list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('list_of_speakers', 'list_of_speakers_id');
CREATE TRIGGER tr_log_structure_level_list_of_speakers_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON structure_level_list_of_speakers_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_point_of_order_category AFTER INSERT OR UPDATE OR DELETE ON point_of_order_category_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('point_of_order_category');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON point_of_order_category_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_point_of_order_category_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON point_of_order_category_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_speaker AFTER INSERT OR UPDATE OR DELETE ON speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('speaker');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON speaker_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_speaker_t_list_of_speakers_id AFTER INSERT OR UPDATE OF list_of_speakers_id OR DELETE ON speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('list_of_speakers', 'list_of_speakers_id');
CREATE TRIGGER tr_log_speaker_t_structure_level_list_of_speakers_id AFTER INSERT OR UPDATE OF structure_level_list_of_speakers_id OR DELETE ON speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('structure_level_list_of_speakers', 'structure_level_list_of_speakers_id');
CREATE TRIGGER tr_log_speaker_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');
CREATE TRIGGER tr_log_speaker_t_point_of_order_category_id AFTER INSERT OR UPDATE OF point_of_order_category_id OR DELETE ON speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('point_of_order_category', 'point_of_order_category_id');
CREATE TRIGGER tr_log_speaker_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_topic AFTER INSERT OR UPDATE OR DELETE ON topic_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('topic');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON topic_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_topic_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON topic_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion AFTER INSERT OR UPDATE OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_t_lead_motion_id AFTER INSERT OR UPDATE OF lead_motion_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'lead_motion_id');
CREATE TRIGGER tr_log_motion_t_sort_parent_id AFTER INSERT OR UPDATE OF sort_parent_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'sort_parent_id');
CREATE TRIGGER tr_log_motion_t_origin_id AFTER INSERT OR UPDATE OF origin_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'origin_id');
CREATE TRIGGER tr_log_motion_t_origin_meeting_id AFTER INSERT OR UPDATE OF origin_meeting_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'origin_meeting_id');

CREATE TRIGGER tr_log_nm_motion_all_derived_motion_ids_motion_t AFTER INSERT OR UPDATE OR DELETE ON nm_motion_all_derived_motion_ids_motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','all_derived_motion_id','motion','all_origin_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_motion_all_derived_motion_ids_motion_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_motion_identical_motion_ids_motion_t AFTER INSERT OR UPDATE OR DELETE ON nm_motion_identical_motion_ids_motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','identical_motion_id_1','motion','identical_motion_id_2');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_motion_identical_motion_ids_motion_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_motion_t_state_id AFTER INSERT OR UPDATE OF state_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_state', 'state_id');
CREATE TRIGGER tr_log_motion_t_recommendation_id AFTER INSERT OR UPDATE OF recommendation_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_state', 'recommendation_id');

CREATE TRIGGER tr_log_state_extension_reference_id_motion_id_gm_motion_state_extension_reference_ids_t AFTER INSERT OR UPDATE OF state_extension_reference_id_motion_id OR DELETE ON gm_motion_state_extension_reference_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','motion_id','motion','state_extension_reference_id_motion_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON gm_motion_state_extension_reference_ids_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_recommendation_extension_reference_id_motion_id_gm_motion_recommendation_extension_reference_ids_t AFTER INSERT OR UPDATE OF recommendation_extension_reference_id_motion_id OR DELETE ON gm_motion_recommendation_extension_reference_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','motion_id','motion','recommendation_extension_reference_id_motion_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON gm_motion_recommendation_extension_reference_ids_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_motion_t_category_id AFTER INSERT OR UPDATE OF category_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_category', 'category_id');
CREATE TRIGGER tr_log_motion_t_block_id AFTER INSERT OR UPDATE OF block_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_block', 'block_id');
CREATE TRIGGER tr_log_motion_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_submitter AFTER INSERT OR UPDATE OR DELETE ON motion_submitter_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_submitter');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_submitter_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_submitter_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON motion_submitter_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');
CREATE TRIGGER tr_log_motion_submitter_t_motion_id AFTER INSERT OR UPDATE OF motion_id OR DELETE ON motion_submitter_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'motion_id');
CREATE TRIGGER tr_log_motion_submitter_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_submitter_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_editor AFTER INSERT OR UPDATE OR DELETE ON motion_editor_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_editor');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_editor_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_editor_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON motion_editor_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');
CREATE TRIGGER tr_log_motion_editor_t_motion_id AFTER INSERT OR UPDATE OF motion_id OR DELETE ON motion_editor_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'motion_id');
CREATE TRIGGER tr_log_motion_editor_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_editor_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_working_group_speaker AFTER INSERT OR UPDATE OR DELETE ON motion_working_group_speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_working_group_speaker');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_working_group_speaker_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_working_group_speaker_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON motion_working_group_speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');
CREATE TRIGGER tr_log_motion_working_group_speaker_t_motion_id AFTER INSERT OR UPDATE OF motion_id OR DELETE ON motion_working_group_speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'motion_id');
CREATE TRIGGER tr_log_motion_working_group_speaker_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_working_group_speaker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_comment AFTER INSERT OR UPDATE OR DELETE ON motion_comment_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_comment');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_comment_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_comment_t_motion_id AFTER INSERT OR UPDATE OF motion_id OR DELETE ON motion_comment_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'motion_id');
CREATE TRIGGER tr_log_motion_comment_t_section_id AFTER INSERT OR UPDATE OF section_id OR DELETE ON motion_comment_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_comment_section', 'section_id');
CREATE TRIGGER tr_log_motion_comment_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_comment_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_comment_section AFTER INSERT OR UPDATE OR DELETE ON motion_comment_section_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_comment_section');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_comment_section_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_comment_section_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_comment_section_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_category AFTER INSERT OR UPDATE OR DELETE ON motion_category_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_category');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_category_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_category_t_parent_id AFTER INSERT OR UPDATE OF parent_id OR DELETE ON motion_category_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_category', 'parent_id');
CREATE TRIGGER tr_log_motion_category_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_category_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_block AFTER INSERT OR UPDATE OR DELETE ON motion_block_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_block');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_block_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_block_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_block_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_change_recommendation AFTER INSERT OR UPDATE OR DELETE ON motion_change_recommendation_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_change_recommendation');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_change_recommendation_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_change_recommendation_t_motion_id AFTER INSERT OR UPDATE OF motion_id OR DELETE ON motion_change_recommendation_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion', 'motion_id');
CREATE TRIGGER tr_log_motion_change_recommendation_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_change_recommendation_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_state AFTER INSERT OR UPDATE OR DELETE ON motion_state_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_state');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_state_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_state_t_submitter_withdraw_state_id AFTER INSERT OR UPDATE OF submitter_withdraw_state_id OR DELETE ON motion_state_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_state', 'submitter_withdraw_state_id');

CREATE TRIGGER tr_log_nm_motion_state_next_state_ids_motion_state_t AFTER INSERT OR UPDATE OR DELETE ON nm_motion_state_next_state_ids_motion_state_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_state','next_state_id','motion_state','previous_state_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_motion_state_next_state_ids_motion_state_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_motion_state_t_workflow_id AFTER INSERT OR UPDATE OF workflow_id OR DELETE ON motion_state_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_workflow', 'workflow_id');
CREATE TRIGGER tr_log_motion_state_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_state_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_motion_workflow AFTER INSERT OR UPDATE OR DELETE ON motion_workflow_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('motion_workflow');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON motion_workflow_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_motion_workflow_t_first_state_id AFTER INSERT OR UPDATE OF first_state_id OR DELETE ON motion_workflow_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_state', 'first_state_id');
CREATE TRIGGER tr_log_motion_workflow_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON motion_workflow_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_poll AFTER INSERT OR UPDATE OR DELETE ON poll_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('poll');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON poll_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_motion_content_object_id_motion_id AFTER INSERT OR UPDATE OF content_object_id_motion_id OR DELETE ON poll_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','content_object_id_motion_id');

CREATE TRIGGER tr_log_assignment_content_object_id_assignment_id AFTER INSERT OR UPDATE OF content_object_id_assignment_id OR DELETE ON poll_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('assignment','content_object_id_assignment_id');

CREATE TRIGGER tr_log_topic_content_object_id_topic_id AFTER INSERT OR UPDATE OF content_object_id_topic_id OR DELETE ON poll_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('topic','content_object_id_topic_id');

CREATE TRIGGER tr_log_nm_poll_voted_ids_user_t AFTER INSERT OR UPDATE OR DELETE ON nm_poll_voted_ids_user_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('poll','poll_id','user','user_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_poll_voted_ids_user_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_poll_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON poll_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_vote AFTER INSERT OR UPDATE OR DELETE ON vote_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('vote');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON vote_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_vote_t_poll_id AFTER INSERT OR UPDATE OF poll_id OR DELETE ON vote_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('poll', 'poll_id');
CREATE TRIGGER tr_log_vote_t_acting_user_id AFTER INSERT OR UPDATE OF acting_user_id OR DELETE ON vote_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('user', 'acting_user_id');
CREATE TRIGGER tr_log_vote_t_represented_user_id AFTER INSERT OR UPDATE OF represented_user_id OR DELETE ON vote_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('user', 'represented_user_id');
CREATE TRIGGER tr_log_vote_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON vote_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_assignment AFTER INSERT OR UPDATE OR DELETE ON assignment_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('assignment');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON assignment_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_assignment_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON assignment_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_assignment_candidate AFTER INSERT OR UPDATE OR DELETE ON assignment_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('assignment_candidate');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON assignment_candidate_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_assignment_candidate_t_assignment_id AFTER INSERT OR UPDATE OF assignment_id OR DELETE ON assignment_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('assignment', 'assignment_id');
CREATE TRIGGER tr_log_assignment_candidate_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON assignment_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');
CREATE TRIGGER tr_log_assignment_candidate_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON assignment_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_poll_candidate_list AFTER INSERT OR UPDATE OR DELETE ON poll_candidate_list_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('poll_candidate_list');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON poll_candidate_list_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_poll_candidate_list_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON poll_candidate_list_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_poll_candidate AFTER INSERT OR UPDATE OR DELETE ON poll_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('poll_candidate');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON poll_candidate_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_poll_candidate_t_poll_candidate_list_id AFTER INSERT OR UPDATE OF poll_candidate_list_id OR DELETE ON poll_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('poll_candidate_list', 'poll_candidate_list_id');
CREATE TRIGGER tr_log_poll_candidate_t_user_id AFTER INSERT OR UPDATE OF user_id OR DELETE ON poll_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('user', 'user_id');
CREATE TRIGGER tr_log_poll_candidate_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON poll_candidate_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_mediafile AFTER INSERT OR UPDATE OR DELETE ON mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('mediafile');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON mediafile_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_mediafile_t_published_to_meetings_in_organization_id AFTER INSERT OR UPDATE OF published_to_meetings_in_organization_id OR DELETE ON mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization', 'published_to_meetings_in_organization_id');
CREATE TRIGGER tr_log_mediafile_t_parent_id AFTER INSERT OR UPDATE OF parent_id OR DELETE ON mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('mediafile', 'parent_id');

CREATE TRIGGER tr_log_meeting_owner_id_meeting_id AFTER INSERT OR UPDATE OF owner_id_meeting_id OR DELETE ON mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting','owner_id_meeting_id');

CREATE TRIGGER tr_log_organization_owner_id_organization_id AFTER INSERT OR UPDATE OF owner_id_organization_id OR DELETE ON mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('organization','owner_id_organization_id');

CREATE TRIGGER tr_log_meeting_mediafile AFTER INSERT OR UPDATE OR DELETE ON meeting_mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('meeting_mediafile');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON meeting_mediafile_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_meeting_mediafile_t_mediafile_id AFTER INSERT OR UPDATE OF mediafile_id OR DELETE ON meeting_mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('mediafile', 'mediafile_id');
CREATE TRIGGER tr_log_meeting_mediafile_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON meeting_mediafile_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_attachment_id_motion_id_gm_meeting_mediafile_attachment_ids_t AFTER INSERT OR UPDATE OF attachment_id_motion_id OR DELETE ON gm_meeting_mediafile_attachment_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile','meeting_mediafile_id','motion','attachment_id_motion_id');

CREATE TRIGGER tr_log_attachment_id_topic_id_gm_meeting_mediafile_attachment_ids_t AFTER INSERT OR UPDATE OF attachment_id_topic_id OR DELETE ON gm_meeting_mediafile_attachment_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile','meeting_mediafile_id','topic','attachment_id_topic_id');

CREATE TRIGGER tr_log_attachment_id_assignment_id_gm_meeting_mediafile_attachment_ids_t AFTER INSERT OR UPDATE OF attachment_id_assignment_id OR DELETE ON gm_meeting_mediafile_attachment_ids_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile','meeting_mediafile_id','assignment','attachment_id_assignment_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON gm_meeting_mediafile_attachment_ids_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_projector AFTER INSERT OR UPDATE OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('projector');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON projector_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_agenda_item_li AFTER INSERT OR UPDATE OF used_as_default_projector_for_agenda_item_list_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_agenda_item_list_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_topic_in_meeti AFTER INSERT OR UPDATE OF used_as_default_projector_for_topic_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_topic_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_list_of_speake AFTER INSERT OR UPDATE OF used_as_default_projector_for_list_of_speakers_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_list_of_speakers_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_current_los_in AFTER INSERT OR UPDATE OF used_as_default_projector_for_current_los_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_current_los_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_motion_in_meet AFTER INSERT OR UPDATE OF used_as_default_projector_for_motion_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_motion_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_amendment_in_m AFTER INSERT OR UPDATE OF used_as_default_projector_for_amendment_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_amendment_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_motion_block_i AFTER INSERT OR UPDATE OF used_as_default_projector_for_motion_block_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_motion_block_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_assignment_in_ AFTER INSERT OR UPDATE OF used_as_default_projector_for_assignment_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_assignment_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_mediafile_in_m AFTER INSERT OR UPDATE OF used_as_default_projector_for_mediafile_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_mediafile_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_message_in_mee AFTER INSERT OR UPDATE OF used_as_default_projector_for_message_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_message_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_countdown_in_m AFTER INSERT OR UPDATE OF used_as_default_projector_for_countdown_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_countdown_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_assignment_pol AFTER INSERT OR UPDATE OF used_as_default_projector_for_assignment_poll_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_assignment_poll_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_motion_poll_in AFTER INSERT OR UPDATE OF used_as_default_projector_for_motion_poll_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_motion_poll_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_used_as_default_projector_for_poll_in_meetin AFTER INSERT OR UPDATE OF used_as_default_projector_for_poll_in_meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'used_as_default_projector_for_poll_in_meeting_id');
CREATE TRIGGER tr_log_projector_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON projector_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_projection AFTER INSERT OR UPDATE OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('projection');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON projection_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_projection_t_current_projector_id AFTER INSERT OR UPDATE OF current_projector_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector', 'current_projector_id');
CREATE TRIGGER tr_log_projection_t_preview_projector_id AFTER INSERT OR UPDATE OF preview_projector_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector', 'preview_projector_id');
CREATE TRIGGER tr_log_projection_t_history_projector_id AFTER INSERT OR UPDATE OF history_projector_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector', 'history_projector_id');

CREATE TRIGGER tr_log_meeting_content_object_id_meeting_id AFTER INSERT OR UPDATE OF content_object_id_meeting_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting','content_object_id_meeting_id');

CREATE TRIGGER tr_log_motion_content_object_id_motion_id AFTER INSERT OR UPDATE OF content_object_id_motion_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion','content_object_id_motion_id');

CREATE TRIGGER tr_log_meeting_mediafile_content_object_id_meeting_mediafile_id AFTER INSERT OR UPDATE OF content_object_id_meeting_mediafile_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_mediafile','content_object_id_meeting_mediafile_id');

CREATE TRIGGER tr_log_list_of_speakers_content_object_id_list_of_speakers_id AFTER INSERT OR UPDATE OF content_object_id_list_of_speakers_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('list_of_speakers','content_object_id_list_of_speakers_id');

CREATE TRIGGER tr_log_motion_block_content_object_id_motion_block_id AFTER INSERT OR UPDATE OF content_object_id_motion_block_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('motion_block','content_object_id_motion_block_id');

CREATE TRIGGER tr_log_assignment_content_object_id_assignment_id AFTER INSERT OR UPDATE OF content_object_id_assignment_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('assignment','content_object_id_assignment_id');

CREATE TRIGGER tr_log_agenda_item_content_object_id_agenda_item_id AFTER INSERT OR UPDATE OF content_object_id_agenda_item_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('agenda_item','content_object_id_agenda_item_id');

CREATE TRIGGER tr_log_topic_content_object_id_topic_id AFTER INSERT OR UPDATE OF content_object_id_topic_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('topic','content_object_id_topic_id');

CREATE TRIGGER tr_log_poll_content_object_id_poll_id AFTER INSERT OR UPDATE OF content_object_id_poll_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('poll','content_object_id_poll_id');

CREATE TRIGGER tr_log_projector_message_content_object_id_projector_message_id AFTER INSERT OR UPDATE OF content_object_id_projector_message_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector_message','content_object_id_projector_message_id');

CREATE TRIGGER tr_log_projector_countdown_content_object_id_projector_countdown_id AFTER INSERT OR UPDATE OF content_object_id_projector_countdown_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('projector_countdown','content_object_id_projector_countdown_id');
CREATE TRIGGER tr_log_projection_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON projection_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_projector_message AFTER INSERT OR UPDATE OR DELETE ON projector_message_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('projector_message');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON projector_message_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_projector_message_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON projector_message_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_projector_countdown AFTER INSERT OR UPDATE OR DELETE ON projector_countdown_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('projector_countdown');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON projector_countdown_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_projector_countdown_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON projector_countdown_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_chat_group AFTER INSERT OR UPDATE OR DELETE ON chat_group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('chat_group');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON chat_group_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


CREATE TRIGGER tr_log_nm_chat_group_read_group_ids_group_t AFTER INSERT OR UPDATE OR DELETE ON nm_chat_group_read_group_ids_group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('chat_group','chat_group_id','group','group_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_chat_group_read_group_ids_group_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_nm_chat_group_write_group_ids_group_t AFTER INSERT OR UPDATE OR DELETE ON nm_chat_group_write_group_ids_group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('chat_group','chat_group_id','group','group_id');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON nm_chat_group_write_group_ids_group_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();
CREATE TRIGGER tr_log_chat_group_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON chat_group_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_chat_message AFTER INSERT OR UPDATE OR DELETE ON chat_message_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('chat_message');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON chat_message_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_chat_message_t_meeting_user_id AFTER INSERT OR UPDATE OF meeting_user_id OR DELETE ON chat_message_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting_user', 'meeting_user_id');
CREATE TRIGGER tr_log_chat_message_t_chat_group_id AFTER INSERT OR UPDATE OF chat_group_id OR DELETE ON chat_message_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('chat_group', 'chat_group_id');
CREATE TRIGGER tr_log_chat_message_t_meeting_id AFTER INSERT OR UPDATE OF meeting_id OR DELETE ON chat_message_t
FOR EACH ROW EXECUTE FUNCTION log_modified_related_models('meeting', 'meeting_id');

CREATE TRIGGER tr_log_action_worker AFTER INSERT OR UPDATE OR DELETE ON action_worker_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('action_worker');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON action_worker_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();

CREATE TRIGGER tr_log_import_preview AFTER INSERT OR UPDATE OR DELETE ON import_preview_t
FOR EACH ROW EXECUTE FUNCTION log_modified_models('import_preview');
CREATE CONSTRAINT TRIGGER notify_transaction_end AFTER INSERT OR UPDATE OR DELETE ON import_preview_t
DEFERRABLE INITIALLY DEFERRED FOR EACH ROW EXECUTE FUNCTION notify_transaction_end();


/*   Relation-list infos
Generated: What will be generated for left field
    FIELD: a usual Database field
    SQL: a sql-expression in a view
    ***: Error
Field Attributes:Field Attributes opposite side
    1: cardinality 1
    1G: cardinality 1 with generic-relation field
    n: cardinality n
    nG: cardinality n with generic-relation-list field
    t: "to" defined
    r: "reference" defined
    s: sql directive inclusive sql-statement
    R: Required
Model.Field -> Model.Field
    model.field names
*/

/*
SQL nr:1rR => organization/gender_ids:-> gender/organization_id
SQL nr:1rR => organization/committee_ids:-> committee/organization_id
SQL nt:1r => organization/active_meeting_ids:-> meeting/is_active_in_organization_id
SQL nt:1r => organization/archived_meeting_ids:-> meeting/is_archived_in_organization_id
SQL nt:1r => organization/template_meeting_ids:-> meeting/template_for_organization_id
SQL nr:1rR => organization/organization_tag_ids:-> organization_tag/organization_id
FIELD 1rR:1t => organization/theme_id:-> theme/theme_for_organization_id
SQL nr:1rR => organization/theme_ids:-> theme/organization_id
SQL nt:1GrR => organization/mediafile_ids:-> mediafile/owner_id
SQL nt:1r => organization/published_mediafile_ids:-> mediafile/published_to_meetings_in_organization_id
SQL nr:1rR => organization/user_ids:-> user/organization_id

FIELD 1r:nr => user/gender_id:-> gender/user_ids
SQL nt:nt => user/is_present_in_meeting_ids:-> meeting/present_user_ids
SQL nts:nts => user/committee_ids:-> committee/user_ids
SQL nt:nt => user/committee_management_ids:-> committee/manager_ids
SQL nt:1rR => user/meeting_user_ids:-> meeting_user/user_id
SQL nt:nt => user/poll_voted_ids:-> poll/voted_ids
SQL nt:1r => user/acting_vote_ids:-> vote/acting_user_id
SQL nt:1r => user/represented_vote_ids:-> vote/represented_user_id
SQL nt:1r => user/poll_candidate_ids:-> poll_candidate/user_id
FIELD 1r:nt => user/home_committee_id:-> committee/native_user_ids
SQL nts:nts => user/meeting_ids:-> meeting/user_ids

FIELD 1rR:nt => meeting_user/user_id:-> user/meeting_user_ids
FIELD 1rR:nt => meeting_user/meeting_id:-> meeting/meeting_user_ids
SQL nt:1rR => meeting_user/personal_note_ids:-> personal_note/meeting_user_id
SQL nt:1r => meeting_user/speaker_ids:-> speaker/meeting_user_id
SQL nt:nt => meeting_user/supported_motion_ids:-> motion/supporter_meeting_user_ids
SQL nt:1rR => meeting_user/motion_editor_ids:-> motion_editor/meeting_user_id
SQL nt:1rR => meeting_user/motion_working_group_speaker_ids:-> motion_working_group_speaker/meeting_user_id
SQL nt:1rR => meeting_user/motion_submitter_ids:-> motion_submitter/meeting_user_id
SQL nt:1r => meeting_user/assignment_candidate_ids:-> assignment_candidate/meeting_user_id
FIELD 1r:nt => meeting_user/vote_delegated_to_id:-> meeting_user/vote_delegations_from_ids
SQL nt:1r => meeting_user/vote_delegations_from_ids:-> meeting_user/vote_delegated_to_id
SQL nt:1r => meeting_user/chat_message_ids:-> chat_message/meeting_user_id
SQL nt:nt => meeting_user/group_ids:-> group/meeting_user_ids
SQL nt:nt => meeting_user/structure_level_ids:-> structure_level/meeting_user_ids

SQL nr:1r => gender/user_ids:-> user/gender_id

SQL nGt:nt,nt => organization_tag/tagged_ids:-> committee/organization_tag_ids,meeting/organization_tag_ids

SQL 1t:1rR => theme/theme_for_organization_id:-> organization/theme_id

SQL nt:1rR => committee/meeting_ids:-> meeting/committee_id
FIELD 1r:1t => committee/default_meeting_id:-> meeting/default_meeting_for_committee_id
SQL nts:nts => committee/user_ids:-> user/committee_ids
SQL nt:nt => committee/manager_ids:-> user/committee_management_ids
FIELD 1r:nt => committee/parent_id:-> committee/child_ids
SQL nt:1r => committee/child_ids:-> committee/parent_id
SQL nt:nt => committee/all_parent_ids:-> committee/all_child_ids
SQL nt:nt => committee/all_child_ids:-> committee/all_parent_ids
SQL nt:1r => committee/native_user_ids:-> user/home_committee_id
SQL nt:nt => committee/forward_to_committee_ids:-> committee/receive_forwardings_from_committee_ids
SQL nt:nt => committee/receive_forwardings_from_committee_ids:-> committee/forward_to_committee_ids
SQL nt:nGt => committee/organization_tag_ids:-> organization_tag/tagged_ids

FIELD 1r:nt => meeting/is_active_in_organization_id:-> organization/active_meeting_ids
FIELD 1r:nt => meeting/is_archived_in_organization_id:-> organization/archived_meeting_ids
FIELD 1r:nt => meeting/template_for_organization_id:-> organization/template_meeting_ids
FIELD 1rR:1t => meeting/motions_default_workflow_id:-> motion_workflow/default_workflow_meeting_id
FIELD 1rR:1t => meeting/motions_default_amendment_workflow_id:-> motion_workflow/default_amendment_workflow_meeting_id
SQL nt:1r => meeting/motion_poll_default_group_ids:-> group/used_as_motion_poll_default_id
SQL nt:1rR => meeting/poll_candidate_list_ids:-> poll_candidate_list/meeting_id
SQL nt:1rR => meeting/poll_candidate_ids:-> poll_candidate/meeting_id
SQL nt:1rR => meeting/meeting_user_ids:-> meeting_user/meeting_id
SQL nt:1r => meeting/assignment_poll_default_group_ids:-> group/used_as_assignment_poll_default_id
SQL nt:1r => meeting/poll_default_group_ids:-> group/used_as_poll_default_id
SQL nt:1r => meeting/topic_poll_default_group_ids:-> group/used_as_topic_poll_default_id
SQL nt:1rR => meeting/projector_ids:-> projector/meeting_id
SQL nt:1rR => meeting/all_projection_ids:-> projection/meeting_id
SQL nt:1rR => meeting/projector_message_ids:-> projector_message/meeting_id
SQL nt:1rR => meeting/projector_countdown_ids:-> projector_countdown/meeting_id
SQL nt:1rR => meeting/tag_ids:-> tag/meeting_id
SQL nt:1rR => meeting/agenda_item_ids:-> agenda_item/meeting_id
SQL nt:1rR => meeting/list_of_speakers_ids:-> list_of_speakers/meeting_id
SQL nt:1rR => meeting/structure_level_list_of_speakers_ids:-> structure_level_list_of_speakers/meeting_id
SQL nt:1rR => meeting/point_of_order_category_ids:-> point_of_order_category/meeting_id
SQL nt:1rR => meeting/speaker_ids:-> speaker/meeting_id
SQL nt:1rR => meeting/topic_ids:-> topic/meeting_id
SQL nt:1rR => meeting/group_ids:-> group/meeting_id
SQL nt:1rR => meeting/meeting_mediafile_ids:-> meeting_mediafile/meeting_id
SQL nt:1GrR => meeting/mediafile_ids:-> mediafile/owner_id
SQL nt:1rR => meeting/motion_ids:-> motion/meeting_id
SQL nt:1r => meeting/forwarded_motion_ids:-> motion/origin_meeting_id
SQL nt:1rR => meeting/motion_comment_section_ids:-> motion_comment_section/meeting_id
SQL nt:1rR => meeting/motion_category_ids:-> motion_category/meeting_id
SQL nt:1rR => meeting/motion_block_ids:-> motion_block/meeting_id
SQL nt:1rR => meeting/motion_workflow_ids:-> motion_workflow/meeting_id
SQL nt:1rR => meeting/motion_comment_ids:-> motion_comment/meeting_id
SQL nt:1rR => meeting/motion_submitter_ids:-> motion_submitter/meeting_id
SQL nt:1rR => meeting/motion_editor_ids:-> motion_editor/meeting_id
SQL nt:1rR => meeting/motion_working_group_speaker_ids:-> motion_working_group_speaker/meeting_id
SQL nt:1rR => meeting/motion_change_recommendation_ids:-> motion_change_recommendation/meeting_id
SQL nt:1rR => meeting/motion_state_ids:-> motion_state/meeting_id
SQL nt:1rR => meeting/poll_ids:-> poll/meeting_id
SQL nt:1rR => meeting/vote_ids:-> vote/meeting_id
SQL nt:1rR => meeting/assignment_ids:-> assignment/meeting_id
SQL nt:1rR => meeting/assignment_candidate_ids:-> assignment_candidate/meeting_id
SQL nt:1rR => meeting/personal_note_ids:-> personal_note/meeting_id
SQL nt:1rR => meeting/chat_group_ids:-> chat_group/meeting_id
SQL nt:1rR => meeting/chat_message_ids:-> chat_message/meeting_id
SQL nt:1rR => meeting/structure_level_ids:-> structure_level/meeting_id
FIELD 1r:1t => meeting/logo_projector_main_id:-> meeting_mediafile/used_as_logo_projector_main_in_meeting_id
FIELD 1r:1t => meeting/logo_projector_header_id:-> meeting_mediafile/used_as_logo_projector_header_in_meeting_id
FIELD 1r:1t => meeting/logo_web_header_id:-> meeting_mediafile/used_as_logo_web_header_in_meeting_id
FIELD 1r:1t => meeting/logo_pdf_header_l_id:-> meeting_mediafile/used_as_logo_pdf_header_l_in_meeting_id
FIELD 1r:1t => meeting/logo_pdf_header_r_id:-> meeting_mediafile/used_as_logo_pdf_header_r_in_meeting_id
FIELD 1r:1t => meeting/logo_pdf_footer_l_id:-> meeting_mediafile/used_as_logo_pdf_footer_l_in_meeting_id
FIELD 1r:1t => meeting/logo_pdf_footer_r_id:-> meeting_mediafile/used_as_logo_pdf_footer_r_in_meeting_id
FIELD 1r:1t => meeting/logo_pdf_ballot_paper_id:-> meeting_mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id
FIELD 1r:1t => meeting/font_regular_id:-> meeting_mediafile/used_as_font_regular_in_meeting_id
FIELD 1r:1t => meeting/font_italic_id:-> meeting_mediafile/used_as_font_italic_in_meeting_id
FIELD 1r:1t => meeting/font_bold_id:-> meeting_mediafile/used_as_font_bold_in_meeting_id
FIELD 1r:1t => meeting/font_bold_italic_id:-> meeting_mediafile/used_as_font_bold_italic_in_meeting_id
FIELD 1r:1t => meeting/font_monospace_id:-> meeting_mediafile/used_as_font_monospace_in_meeting_id
FIELD 1r:1t => meeting/font_chyron_speaker_name_id:-> meeting_mediafile/used_as_font_chyron_speaker_name_in_meeting_id
FIELD 1r:1t => meeting/font_projector_h1_id:-> meeting_mediafile/used_as_font_projector_h1_in_meeting_id
FIELD 1r:1t => meeting/font_projector_h2_id:-> meeting_mediafile/used_as_font_projector_h2_in_meeting_id
FIELD 1rR:nt => meeting/committee_id:-> committee/meeting_ids
SQL 1t:1r => meeting/default_meeting_for_committee_id:-> committee/default_meeting_id
SQL nt:nGt => meeting/organization_tag_ids:-> organization_tag/tagged_ids
SQL nt:nt => meeting/present_user_ids:-> user/is_present_in_meeting_ids
SQL nts:nts => meeting/user_ids:-> user/meeting_ids
FIELD 1rR:1t => meeting/reference_projector_id:-> projector/used_as_reference_projector_meeting_id
FIELD 1r:1t => meeting/list_of_speakers_countdown_id:-> projector_countdown/used_as_list_of_speakers_countdown_meeting_id
FIELD 1r:1t => meeting/poll_countdown_id:-> projector_countdown/used_as_poll_countdown_meeting_id
SQL nt:1GrR => meeting/projection_ids:-> projection/content_object_id
SQL ntR:1r => meeting/default_projector_agenda_item_list_ids:-> projector/used_as_default_projector_for_agenda_item_list_in_meeting_id
SQL ntR:1r => meeting/default_projector_topic_ids:-> projector/used_as_default_projector_for_topic_in_meeting_id
SQL ntR:1r => meeting/default_projector_list_of_speakers_ids:-> projector/used_as_default_projector_for_list_of_speakers_in_meeting_id
SQL ntR:1r => meeting/default_projector_current_los_ids:-> projector/used_as_default_projector_for_current_los_in_meeting_id
SQL ntR:1r => meeting/default_projector_motion_ids:-> projector/used_as_default_projector_for_motion_in_meeting_id
SQL ntR:1r => meeting/default_projector_amendment_ids:-> projector/used_as_default_projector_for_amendment_in_meeting_id
SQL ntR:1r => meeting/default_projector_motion_block_ids:-> projector/used_as_default_projector_for_motion_block_in_meeting_id
SQL ntR:1r => meeting/default_projector_assignment_ids:-> projector/used_as_default_projector_for_assignment_in_meeting_id
SQL ntR:1r => meeting/default_projector_mediafile_ids:-> projector/used_as_default_projector_for_mediafile_in_meeting_id
SQL ntR:1r => meeting/default_projector_message_ids:-> projector/used_as_default_projector_for_message_in_meeting_id
SQL ntR:1r => meeting/default_projector_countdown_ids:-> projector/used_as_default_projector_for_countdown_in_meeting_id
SQL ntR:1r => meeting/default_projector_assignment_poll_ids:-> projector/used_as_default_projector_for_assignment_poll_in_meeting_id
SQL ntR:1r => meeting/default_projector_motion_poll_ids:-> projector/used_as_default_projector_for_motion_poll_in_meeting_id
SQL ntR:1r => meeting/default_projector_poll_ids:-> projector/used_as_default_projector_for_poll_in_meeting_id
FIELD 1rR:1t => meeting/default_group_id:-> group/default_group_for_meeting_id
FIELD 1r:1t => meeting/admin_group_id:-> group/admin_group_for_meeting_id
FIELD 1r:1t => meeting/anonymous_group_id:-> group/anonymous_group_for_meeting_id

SQL nt:nt => structure_level/meeting_user_ids:-> meeting_user/structure_level_ids
SQL nt:1rR => structure_level/structure_level_list_of_speakers_ids:-> structure_level_list_of_speakers/structure_level_id
FIELD 1rR:nt => structure_level/meeting_id:-> meeting/structure_level_ids

SQL nt:nt => group/meeting_user_ids:-> meeting_user/group_ids
SQL 1t:1rR => group/default_group_for_meeting_id:-> meeting/default_group_id
SQL 1t:1r => group/admin_group_for_meeting_id:-> meeting/admin_group_id
SQL 1t:1r => group/anonymous_group_for_meeting_id:-> meeting/anonymous_group_id
SQL nt:nt => group/meeting_mediafile_access_group_ids:-> meeting_mediafile/access_group_ids
SQL nt:nt => group/meeting_mediafile_inherited_access_group_ids:-> meeting_mediafile/inherited_access_group_ids
SQL nt:nt => group/read_comment_section_ids:-> motion_comment_section/read_group_ids
SQL nt:nt => group/write_comment_section_ids:-> motion_comment_section/write_group_ids
SQL nt:nt => group/read_chat_group_ids:-> chat_group/read_group_ids
SQL nt:nt => group/write_chat_group_ids:-> chat_group/write_group_ids
SQL nt:nt => group/poll_ids:-> poll/entitled_group_ids
FIELD 1r:nt => group/used_as_motion_poll_default_id:-> meeting/motion_poll_default_group_ids
FIELD 1r:nt => group/used_as_assignment_poll_default_id:-> meeting/assignment_poll_default_group_ids
FIELD 1r:nt => group/used_as_topic_poll_default_id:-> meeting/topic_poll_default_group_ids
FIELD 1r:nt => group/used_as_poll_default_id:-> meeting/poll_default_group_ids
FIELD 1rR:nt => group/meeting_id:-> meeting/group_ids

FIELD 1rR:nt => personal_note/meeting_user_id:-> meeting_user/personal_note_ids
FIELD 1Gr:nt => personal_note/content_object_id:-> motion/personal_note_ids
FIELD 1rR:nt => personal_note/meeting_id:-> meeting/personal_note_ids

SQL nGt:nt,nt,nt => tag/tagged_ids:-> agenda_item/tag_ids,assignment/tag_ids,motion/tag_ids
FIELD 1rR:nt => tag/meeting_id:-> meeting/tag_ids

FIELD 1GrR:1t,1t,1t,1tR => agenda_item/content_object_id:-> motion/agenda_item_id,motion_block/agenda_item_id,assignment/agenda_item_id,topic/agenda_item_id
FIELD 1r:nt => agenda_item/parent_id:-> agenda_item/child_ids
SQL nt:1r => agenda_item/child_ids:-> agenda_item/parent_id
SQL nt:nGt => agenda_item/tag_ids:-> tag/tagged_ids
SQL nt:1GrR => agenda_item/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => agenda_item/meeting_id:-> meeting/agenda_item_ids

FIELD 1GrR:1tR,1tR,1tR,1tR,1t => list_of_speakers/content_object_id:-> motion/list_of_speakers_id,motion_block/list_of_speakers_id,assignment/list_of_speakers_id,topic/list_of_speakers_id,meeting_mediafile/list_of_speakers_id
SQL nt:1rR => list_of_speakers/speaker_ids:-> speaker/list_of_speakers_id
SQL nt:1rR => list_of_speakers/structure_level_list_of_speakers_ids:-> structure_level_list_of_speakers/list_of_speakers_id
SQL nt:1GrR => list_of_speakers/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => list_of_speakers/meeting_id:-> meeting/list_of_speakers_ids

FIELD 1rR:nt => structure_level_list_of_speakers/structure_level_id:-> structure_level/structure_level_list_of_speakers_ids
FIELD 1rR:nt => structure_level_list_of_speakers/list_of_speakers_id:-> list_of_speakers/structure_level_list_of_speakers_ids
SQL nt:1r => structure_level_list_of_speakers/speaker_ids:-> speaker/structure_level_list_of_speakers_id
FIELD 1rR:nt => structure_level_list_of_speakers/meeting_id:-> meeting/structure_level_list_of_speakers_ids

FIELD 1rR:nt => point_of_order_category/meeting_id:-> meeting/point_of_order_category_ids
SQL nt:1r => point_of_order_category/speaker_ids:-> speaker/point_of_order_category_id

FIELD 1rR:nt => speaker/list_of_speakers_id:-> list_of_speakers/speaker_ids
FIELD 1r:nt => speaker/structure_level_list_of_speakers_id:-> structure_level_list_of_speakers/speaker_ids
FIELD 1r:nt => speaker/meeting_user_id:-> meeting_user/speaker_ids
FIELD 1r:nt => speaker/point_of_order_category_id:-> point_of_order_category/speaker_ids
FIELD 1rR:nt => speaker/meeting_id:-> meeting/speaker_ids

SQL nt:nGt => topic/attachment_meeting_mediafile_ids:-> meeting_mediafile/attachment_ids
SQL 1tR:1GrR => topic/agenda_item_id:-> agenda_item/content_object_id
SQL 1tR:1GrR => topic/list_of_speakers_id:-> list_of_speakers/content_object_id
SQL nt:1GrR => topic/poll_ids:-> poll/content_object_id
SQL nt:1GrR => topic/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => topic/meeting_id:-> meeting/topic_ids

FIELD 1r:nt => motion/lead_motion_id:-> motion/amendment_ids
SQL nt:1r => motion/amendment_ids:-> motion/lead_motion_id
FIELD 1r:nt => motion/sort_parent_id:-> motion/sort_child_ids
SQL nt:1r => motion/sort_child_ids:-> motion/sort_parent_id
FIELD 1r:nt => motion/origin_id:-> motion/derived_motion_ids
FIELD 1r:nt => motion/origin_meeting_id:-> meeting/forwarded_motion_ids
SQL nt:1r => motion/derived_motion_ids:-> motion/origin_id
SQL nt:nt => motion/all_origin_ids:-> motion/all_derived_motion_ids
SQL nt:nt => motion/all_derived_motion_ids:-> motion/all_origin_ids
SQL nt:nt => motion/identical_motion_ids:-> motion/identical_motion_ids
FIELD 1rR:nt => motion/state_id:-> motion_state/motion_ids
FIELD 1r:nt => motion/recommendation_id:-> motion_state/motion_recommendation_ids
SQL nGt:nt => motion/state_extension_reference_ids:-> motion/referenced_in_motion_state_extension_ids
SQL nt:nGt => motion/referenced_in_motion_state_extension_ids:-> motion/state_extension_reference_ids
SQL nGt:nt => motion/recommendation_extension_reference_ids:-> motion/referenced_in_motion_recommendation_extension_ids
SQL nt:nGt => motion/referenced_in_motion_recommendation_extension_ids:-> motion/recommendation_extension_reference_ids
FIELD 1r:nt => motion/category_id:-> motion_category/motion_ids
FIELD 1r:nt => motion/block_id:-> motion_block/motion_ids
SQL nt:1rR => motion/submitter_ids:-> motion_submitter/motion_id
SQL nt:nt => motion/supporter_meeting_user_ids:-> meeting_user/supported_motion_ids
SQL nt:1rR => motion/editor_ids:-> motion_editor/motion_id
SQL nt:1rR => motion/working_group_speaker_ids:-> motion_working_group_speaker/motion_id
SQL nt:1GrR => motion/poll_ids:-> poll/content_object_id
SQL nt:1rR => motion/change_recommendation_ids:-> motion_change_recommendation/motion_id
SQL nt:1rR => motion/comment_ids:-> motion_comment/motion_id
SQL 1t:1GrR => motion/agenda_item_id:-> agenda_item/content_object_id
SQL 1tR:1GrR => motion/list_of_speakers_id:-> list_of_speakers/content_object_id
SQL nt:nGt => motion/tag_ids:-> tag/tagged_ids
SQL nt:nGt => motion/attachment_meeting_mediafile_ids:-> meeting_mediafile/attachment_ids
SQL nt:1GrR => motion/projection_ids:-> projection/content_object_id
SQL nt:1Gr => motion/personal_note_ids:-> personal_note/content_object_id
FIELD 1rR:nt => motion/meeting_id:-> meeting/motion_ids

FIELD 1rR:nt => motion_submitter/meeting_user_id:-> meeting_user/motion_submitter_ids
FIELD 1rR:nt => motion_submitter/motion_id:-> motion/submitter_ids
FIELD 1rR:nt => motion_submitter/meeting_id:-> meeting/motion_submitter_ids

FIELD 1rR:nt => motion_editor/meeting_user_id:-> meeting_user/motion_editor_ids
FIELD 1rR:nt => motion_editor/motion_id:-> motion/editor_ids
FIELD 1rR:nt => motion_editor/meeting_id:-> meeting/motion_editor_ids

FIELD 1rR:nt => motion_working_group_speaker/meeting_user_id:-> meeting_user/motion_working_group_speaker_ids
FIELD 1rR:nt => motion_working_group_speaker/motion_id:-> motion/working_group_speaker_ids
FIELD 1rR:nt => motion_working_group_speaker/meeting_id:-> meeting/motion_working_group_speaker_ids

FIELD 1rR:nt => motion_comment/motion_id:-> motion/comment_ids
FIELD 1rR:nt => motion_comment/section_id:-> motion_comment_section/comment_ids
FIELD 1rR:nt => motion_comment/meeting_id:-> meeting/motion_comment_ids

SQL nt:1rR => motion_comment_section/comment_ids:-> motion_comment/section_id
SQL nt:nt => motion_comment_section/read_group_ids:-> group/read_comment_section_ids
SQL nt:nt => motion_comment_section/write_group_ids:-> group/write_comment_section_ids
FIELD 1rR:nt => motion_comment_section/meeting_id:-> meeting/motion_comment_section_ids

FIELD 1r:nt => motion_category/parent_id:-> motion_category/child_ids
SQL nt:1r => motion_category/child_ids:-> motion_category/parent_id
SQL nt:1r => motion_category/motion_ids:-> motion/category_id
FIELD 1rR:nt => motion_category/meeting_id:-> meeting/motion_category_ids

SQL nt:1r => motion_block/motion_ids:-> motion/block_id
SQL 1t:1GrR => motion_block/agenda_item_id:-> agenda_item/content_object_id
SQL 1tR:1GrR => motion_block/list_of_speakers_id:-> list_of_speakers/content_object_id
SQL nt:1GrR => motion_block/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => motion_block/meeting_id:-> meeting/motion_block_ids

FIELD 1rR:nt => motion_change_recommendation/motion_id:-> motion/change_recommendation_ids
FIELD 1rR:nt => motion_change_recommendation/meeting_id:-> meeting/motion_change_recommendation_ids

FIELD 1r:nt => motion_state/submitter_withdraw_state_id:-> motion_state/submitter_withdraw_back_ids
SQL nt:1r => motion_state/submitter_withdraw_back_ids:-> motion_state/submitter_withdraw_state_id
SQL nt:nt => motion_state/next_state_ids:-> motion_state/previous_state_ids
SQL nt:nt => motion_state/previous_state_ids:-> motion_state/next_state_ids
SQL nt:1rR => motion_state/motion_ids:-> motion/state_id
SQL nt:1r => motion_state/motion_recommendation_ids:-> motion/recommendation_id
FIELD 1rR:nt => motion_state/workflow_id:-> motion_workflow/state_ids
SQL 1t:1rR => motion_state/first_state_of_workflow_id:-> motion_workflow/first_state_id
FIELD 1rR:nt => motion_state/meeting_id:-> meeting/motion_state_ids

SQL nt:1rR => motion_workflow/state_ids:-> motion_state/workflow_id
FIELD 1rR:1t => motion_workflow/first_state_id:-> motion_state/first_state_of_workflow_id
SQL 1t:1rR => motion_workflow/default_workflow_meeting_id:-> meeting/motions_default_workflow_id
SQL 1t:1rR => motion_workflow/default_amendment_workflow_meeting_id:-> meeting/motions_default_amendment_workflow_id
FIELD 1rR:nt => motion_workflow/meeting_id:-> meeting/motion_workflow_ids

FIELD 1GrR:nt,nt,nt => poll/content_object_id:-> motion/poll_ids,assignment/poll_ids,topic/poll_ids
SQL nt:1rR => poll/vote_ids:-> vote/poll_id
SQL nt:nt => poll/voted_ids:-> user/poll_voted_ids
SQL nt:nt => poll/entitled_group_ids:-> group/poll_ids
SQL nt:1GrR => poll/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => poll/meeting_id:-> meeting/poll_ids

FIELD 1rR:nt => vote/poll_id:-> poll/vote_ids
FIELD 1r:nt => vote/acting_user_id:-> user/acting_vote_ids
FIELD 1r:nt => vote/represented_user_id:-> user/represented_vote_ids
FIELD 1rR:nt => vote/meeting_id:-> meeting/vote_ids

SQL nt:1rR => assignment/candidate_ids:-> assignment_candidate/assignment_id
SQL nt:1GrR => assignment/poll_ids:-> poll/content_object_id
SQL 1t:1GrR => assignment/agenda_item_id:-> agenda_item/content_object_id
SQL 1tR:1GrR => assignment/list_of_speakers_id:-> list_of_speakers/content_object_id
SQL nt:nGt => assignment/tag_ids:-> tag/tagged_ids
SQL nt:nGt => assignment/attachment_meeting_mediafile_ids:-> meeting_mediafile/attachment_ids
SQL nt:1GrR => assignment/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => assignment/meeting_id:-> meeting/assignment_ids

FIELD 1rR:nt => assignment_candidate/assignment_id:-> assignment/candidate_ids
FIELD 1r:nt => assignment_candidate/meeting_user_id:-> meeting_user/assignment_candidate_ids
FIELD 1rR:nt => assignment_candidate/meeting_id:-> meeting/assignment_candidate_ids

SQL nt:1rR => poll_candidate_list/poll_candidate_ids:-> poll_candidate/poll_candidate_list_id
FIELD 1rR:nt => poll_candidate_list/meeting_id:-> meeting/poll_candidate_list_ids

FIELD 1rR:nt => poll_candidate/poll_candidate_list_id:-> poll_candidate_list/poll_candidate_ids
FIELD 1r:nt => poll_candidate/user_id:-> user/poll_candidate_ids
FIELD 1rR:nt => poll_candidate/meeting_id:-> meeting/poll_candidate_ids

FIELD 1r:nt => mediafile/published_to_meetings_in_organization_id:-> organization/published_mediafile_ids
FIELD 1r:nt => mediafile/parent_id:-> mediafile/child_ids
SQL nt:1r => mediafile/child_ids:-> mediafile/parent_id
FIELD 1GrR:nt,nt => mediafile/owner_id:-> meeting/mediafile_ids,organization/mediafile_ids
SQL nt:1rR => mediafile/meeting_mediafile_ids:-> meeting_mediafile/mediafile_id

FIELD 1rR:nt => meeting_mediafile/mediafile_id:-> mediafile/meeting_mediafile_ids
FIELD 1rR:nt => meeting_mediafile/meeting_id:-> meeting/meeting_mediafile_ids
SQL nt:nt => meeting_mediafile/inherited_access_group_ids:-> group/meeting_mediafile_inherited_access_group_ids
SQL nt:nt => meeting_mediafile/access_group_ids:-> group/meeting_mediafile_access_group_ids
SQL 1t:1GrR => meeting_mediafile/list_of_speakers_id:-> list_of_speakers/content_object_id
SQL nt:1GrR => meeting_mediafile/projection_ids:-> projection/content_object_id
SQL nGt:nt,nt,nt => meeting_mediafile/attachment_ids:-> motion/attachment_meeting_mediafile_ids,topic/attachment_meeting_mediafile_ids,assignment/attachment_meeting_mediafile_ids
SQL 1t:1r => meeting_mediafile/used_as_logo_projector_main_in_meeting_id:-> meeting/logo_projector_main_id
SQL 1t:1r => meeting_mediafile/used_as_logo_projector_header_in_meeting_id:-> meeting/logo_projector_header_id
SQL 1t:1r => meeting_mediafile/used_as_logo_web_header_in_meeting_id:-> meeting/logo_web_header_id
SQL 1t:1r => meeting_mediafile/used_as_logo_pdf_header_l_in_meeting_id:-> meeting/logo_pdf_header_l_id
SQL 1t:1r => meeting_mediafile/used_as_logo_pdf_header_r_in_meeting_id:-> meeting/logo_pdf_header_r_id
SQL 1t:1r => meeting_mediafile/used_as_logo_pdf_footer_l_in_meeting_id:-> meeting/logo_pdf_footer_l_id
SQL 1t:1r => meeting_mediafile/used_as_logo_pdf_footer_r_in_meeting_id:-> meeting/logo_pdf_footer_r_id
SQL 1t:1r => meeting_mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id:-> meeting/logo_pdf_ballot_paper_id
SQL 1t:1r => meeting_mediafile/used_as_font_regular_in_meeting_id:-> meeting/font_regular_id
SQL 1t:1r => meeting_mediafile/used_as_font_italic_in_meeting_id:-> meeting/font_italic_id
SQL 1t:1r => meeting_mediafile/used_as_font_bold_in_meeting_id:-> meeting/font_bold_id
SQL 1t:1r => meeting_mediafile/used_as_font_bold_italic_in_meeting_id:-> meeting/font_bold_italic_id
SQL 1t:1r => meeting_mediafile/used_as_font_monospace_in_meeting_id:-> meeting/font_monospace_id
SQL 1t:1r => meeting_mediafile/used_as_font_chyron_speaker_name_in_meeting_id:-> meeting/font_chyron_speaker_name_id
SQL 1t:1r => meeting_mediafile/used_as_font_projector_h1_in_meeting_id:-> meeting/font_projector_h1_id
SQL 1t:1r => meeting_mediafile/used_as_font_projector_h2_in_meeting_id:-> meeting/font_projector_h2_id

SQL nt:1r => projector/current_projection_ids:-> projection/current_projector_id
SQL nt:1r => projector/preview_projection_ids:-> projection/preview_projector_id
SQL nt:1r => projector/history_projection_ids:-> projection/history_projector_id
SQL 1t:1rR => projector/used_as_reference_projector_meeting_id:-> meeting/reference_projector_id
FIELD 1r:ntR => projector/used_as_default_projector_for_agenda_item_list_in_meeting_id:-> meeting/default_projector_agenda_item_list_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_topic_in_meeting_id:-> meeting/default_projector_topic_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_list_of_speakers_in_meeting_id:-> meeting/default_projector_list_of_speakers_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_current_los_in_meeting_id:-> meeting/default_projector_current_los_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_motion_in_meeting_id:-> meeting/default_projector_motion_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_amendment_in_meeting_id:-> meeting/default_projector_amendment_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_motion_block_in_meeting_id:-> meeting/default_projector_motion_block_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_assignment_in_meeting_id:-> meeting/default_projector_assignment_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_mediafile_in_meeting_id:-> meeting/default_projector_mediafile_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_message_in_meeting_id:-> meeting/default_projector_message_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_countdown_in_meeting_id:-> meeting/default_projector_countdown_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_assignment_poll_in_meeting_id:-> meeting/default_projector_assignment_poll_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_motion_poll_in_meeting_id:-> meeting/default_projector_motion_poll_ids
FIELD 1r:ntR => projector/used_as_default_projector_for_poll_in_meeting_id:-> meeting/default_projector_poll_ids
FIELD 1rR:nt => projector/meeting_id:-> meeting/projector_ids

FIELD 1r:nt => projection/current_projector_id:-> projector/current_projection_ids
FIELD 1r:nt => projection/preview_projector_id:-> projector/preview_projection_ids
FIELD 1r:nt => projection/history_projector_id:-> projector/history_projection_ids
FIELD 1GrR:nt,nt,nt,nt,nt,nt,nt,nt,nt,nt,nt => projection/content_object_id:-> meeting/projection_ids,motion/projection_ids,meeting_mediafile/projection_ids,list_of_speakers/projection_ids,motion_block/projection_ids,assignment/projection_ids,agenda_item/projection_ids,topic/projection_ids,poll/projection_ids,projector_message/projection_ids,projector_countdown/projection_ids
FIELD 1rR:nt => projection/meeting_id:-> meeting/all_projection_ids

SQL nt:1GrR => projector_message/projection_ids:-> projection/content_object_id
FIELD 1rR:nt => projector_message/meeting_id:-> meeting/projector_message_ids

SQL nt:1GrR => projector_countdown/projection_ids:-> projection/content_object_id
SQL 1t:1r => projector_countdown/used_as_list_of_speakers_countdown_meeting_id:-> meeting/list_of_speakers_countdown_id
SQL 1t:1r => projector_countdown/used_as_poll_countdown_meeting_id:-> meeting/poll_countdown_id
FIELD 1rR:nt => projector_countdown/meeting_id:-> meeting/projector_countdown_ids

SQL nt:1rR => chat_group/chat_message_ids:-> chat_message/chat_group_id
SQL nt:nt => chat_group/read_group_ids:-> group/read_chat_group_ids
SQL nt:nt => chat_group/write_group_ids:-> group/write_chat_group_ids
FIELD 1rR:nt => chat_group/meeting_id:-> meeting/chat_group_ids

FIELD 1r:nt => chat_message/meeting_user_id:-> meeting_user/chat_message_ids
FIELD 1rR:nt => chat_message/chat_group_id:-> chat_group/chat_message_ids
FIELD 1rR:nt => chat_message/meeting_id:-> meeting/chat_message_ids

*/
/*
There are 1 errors/warnings
    projection/content: type:JSON is marked as a calculated field and not generated in schema
*/

/*   Missing attribute handling for constant, on_delete, equal_fields, unique, deferred */