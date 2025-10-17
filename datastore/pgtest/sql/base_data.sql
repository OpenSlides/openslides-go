-- This script can only be used for an empty database without used sequences.
INSERT INTO
theme_t (id, name)
VALUES
(1, 'standard theme');
-- Increase sequence of theme_t.id to avoid errors.
SELECT nextval('theme_t_id_seq');

INSERT INTO
organization_t (name, theme_id)
VALUES
('Intevation', 1);

INSERT INTO
committee_t (id, name)
VALUES
(1, 'committee');
-- Increase sequence of committee_t.id to avoid errors.
SELECT nextval('committee_t_id_seq');

INSERT INTO gender_t (name) VALUES ('female');


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
(1, 1, 2, 1, 1, 1, 1, 'meeting');
-- Increase sequence of committee_t.id to avoid errors.
SELECT nextval('meeting_t_id_seq');

INSERT INTO
group_t (id, name, meeting_id, permissions)
VALUES
(
    1,
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
(2, 'Admin', 1, DEFAULT);
-- Set sequence of group_t.id to avoid errors.
SELECT setval('group_t_id_seq', 2);

INSERT INTO
motion_workflow_t (
    id,
    name,
    sequential_number,
    first_state_id,
    meeting_id
)
VALUES
(1, 'Simple Workflow', 1, 1, 1);
-- Increase sequence of motion_workflow_t.id to avoid errors.
SELECT nextval('motion_workflow_t_id_seq');

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
    1,
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
    2,
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
    3,
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
    4,
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
-- Set sequence of motion_state_t.id to avoid errors.
SELECT setval('motion_state_t_id_seq', 4);


INSERT INTO
projector_t
(
    id,
    sequential_number,
    meeting_id,
    used_as_default_projector_for_agenda_item_list_in_meeting_id,
    used_as_default_projector_for_topic_in_meeting_id,
    used_as_default_projector_for_list_of_speakers_in_meeting_id,
    used_as_default_projector_for_current_los_in_meeting_id,
    used_as_default_projector_for_motion_in_meeting_id,
    used_as_default_projector_for_amendment_in_meeting_id,
    used_as_default_projector_for_motion_block_in_meeting_id,
    used_as_default_projector_for_assignment_in_meeting_id,
    used_as_default_projector_for_mediafile_in_meeting_id,
    used_as_default_projector_for_message_in_meeting_id,
    used_as_default_projector_for_countdown_in_meeting_id,
    used_as_default_projector_for_assignment_poll_in_meeting_id,
    used_as_default_projector_for_motion_poll_in_meeting_id,
    used_as_default_projector_for_poll_in_meeting_id
)
VALUES (1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1);

COMMIT;
