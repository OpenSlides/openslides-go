-- This script can only be used for an empty database without used sequences.
INSERT INTO gender_t (name) VALUES('female');
INSERT INTO user_t (username,gender_id) VALUES('tom',1); --relation --relation-list gender_ids
BEGIN;
    INSERT INTO motion_state_t (id,name,weight,workflow_id,meeting_id)
        VALUES (1,'motionState1',1,1,1);
    INSERT INTO motion_workflow_t (id, name, sequential_number, first_state_id, meeting_id)
        VALUES (1,'workflow1',1,1,1);
    INSERT INTO meeting_t (id,name,motions_default_workflow_id,motions_default_amendment_workflow_id,committee_id,reference_projector_id,default_group_id)
        VALUES (1,'name',1,1,1,1,1);
    INSERT INTO organization_tag_t(id,name,color)
        VALUES(1,'tagA','#cc3b03'); --generic-relation-list tagged_ids
    Insert Into committee(id, name, default_meeting_id)
        Values(1,'plenum',1); --relation-list organization_tag_ids --relation 1:1 default_meeting_id
    INSERT INTO projector_t (id,sequential_number,meeting_id)
        VALUES (1,1,1);
    INSERT INTO group_t (id,name,meeting_id)
        VALUES (1,'gruppe1',1);
COMMIT;

INSERT INTO organization_tag_t (id,name,color)
    VALUES (2,'bunt','#ffffff');
INSERT INTO gm_organization_tag_tagged_ids_t(organization_tag_id,tagged_id)
    VALUES(2,'meeting/1'); 

INSERT INTO topic_t (id,title, sequential_number, meeting_id)
    VALUES (1, 'Thema3', 1, 1);
INSERT INTO agenda_item_t (content_object_id, meeting_id)
    VALUES ('topic/1',1);--agenda_item.content_object_id:topic.agenda_item_id gr:r

INSERT INTO option_t (id,content_object_id,meeting_id)
    VALUES (1, 'user/1', 1);--rl:gr user.option_id:option.content_object_id

INSERT INTO nm_committee_manager_ids_user_t
    VALUES (1,1); --rl:rl committee_ids:user_ids

INSERT INTO
    theme_t (id, name)
VALUES
    (1, 'standard theme');

INSERT INTO
    organization_t (name, theme_id)
VALUES
    ('Intevation', 1);

INSERT INTO
    committee_t (id, name)
VALUES
    (2, 'committee');

BEGIN;
    INSERT INTO
        meeting_t (
            id,
            default_group_id,
            admin_group_id,
            motions_default_workflow_id,
            motions_default_amendment_workflow_id,
            committee_id,
            reference_projector_id,
            name
        )
    VALUES
        (2, 2, 2, 2, 2, 1, 2, 'meeting');

    INSERT INTO
        group_t (id, name, meeting_id, permissions)
    VALUES
        (
            2,
            'Default',
            1,
            '{
                "agenda_item.can_see",
                "assignment.can_see",
                "meeting.can_see_autopilot",
                "meeting.can_see_frontpage",
                "motion.can_see",
                "projector.can_see"
            }'
        ),
        (3, 'Admin', 1, DEFAULT);

    INSERT INTO
        motion_workflow_t (
            id,
            name,
            sequential_number,
            first_state_id,
            meeting_id
        )
    VALUES
        (2, 'Simple Workflow', 1, 2, 2);

    INSERT INTO
        motion_state_t (
            id,
            name,
            weight,
            workflow_id,
            meeting_id,
            allow_create_poll,
            allow_support,
            set_workflow_timestamp,
            recommendation_label,
            css_class,
            merge_amendment_into_final
        )
    VALUES
        (
            2,
            'submitted',
            1,
            1,
            1,
            true,
            true,
            true,
            'Submitted',
            'grey',
            'do_not_merge'
        ),
        (
            3,
            'accepted',
            2,
            1,
            1,
            DEFAULT,
            DEFAULT,
            DEFAULT,
            'Acceptance',
            'green',
            'do_merge'
        ),
        (
            4,
            'rejected',
            3,
            1,
            1,
            DEFAULT,
            DEFAULT,
            DEFAULT,
            'Rejection',
            'red',
            'do_not_merge'
        ),
        (
            5,
            'not decided',
            4,
            1,
            1,
            DEFAULT,
            DEFAULT,
            DEFAULT,
            'No decision',
            'grey',
            'do_not_merge'
        );

    INSERT INTO projector_t (id,sequential_number,meeting_id) 
        VALUES (2,2,1);

COMMIT;
