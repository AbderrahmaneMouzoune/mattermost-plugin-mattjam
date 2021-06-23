import {connect} from 'react-redux';
import {bindActionCreators, Dispatch, ActionCreatorsMapObject} from 'redux';

import {getBool, getTheme} from 'mattermost-redux/selectors/entities/preferences';
import {GenericAction, ActionFunc, ActionResult} from 'mattermost-redux/types/actions';

import {Post} from 'mattermost-redux/types/posts';

import {GlobalState} from '../../types';
import {displayUsernameForUser} from '../../utils/user_utils';
import {enrichMeetingJwt, openMattJamMeeting} from '../../actions';

import {PostTypeMattJam} from './post_type_mattjam';

type OwnProps = {
    post: Post,
}

function mapStateToProps(state: GlobalState, ownProps: OwnProps) {
    const post = ownProps.post;
    const user = state.entities.users.profiles[post.user_id];
    const config = state['plugins-mattjam'].config;

    return {
        ...ownProps,
        theme: getTheme(state),
        creatorName: displayUsernameForUser(user, state.entities.general.config),
        useMilitaryTime: getBool(state, 'display_settings', 'use_military_time', false),
        meetingEmbedded: Boolean(config.embedded)
    };
}

type Actions = {
    enrichMeetingJwt: (jwt: string) => Promise<ActionResult>,
    openMattJamMeeting: (post: Post | null, jwt: string | null) => ActionResult,
}

function mapDispatchToProps(dispatch: Dispatch<GenericAction>) {
    return {
        actions: bindActionCreators<ActionCreatorsMapObject<ActionFunc>, Actions>({
            enrichMeetingJwt,
            openMattJamMeeting
        }, dispatch)
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(PostTypeMattJam);
