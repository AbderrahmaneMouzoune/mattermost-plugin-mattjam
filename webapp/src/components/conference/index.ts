import {connect} from 'react-redux';
import {bindActionCreators, Dispatch} from 'redux';

import {GenericAction} from 'mattermost-redux/types/actions';

import {GlobalState} from '../../types';
import Conference from './conference';
import {openMattJamMeeting} from '../../actions';

function mapStateToProps(state: GlobalState) {
    return {
        post: state['plugins-mattjam'].openMeeting,
        jwt: state['plugins-mattjam'].openMeetingJwt
    };
}

function mapDispatchToProps(dispatch: Dispatch<GenericAction>) {
    return {
        actions: bindActionCreators({
            openMattJamMeeting
        }, dispatch)
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(Conference);
