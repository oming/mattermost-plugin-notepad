import {combineReducers} from 'redux';

import {RECEIVED_BOOKMARK_SETTINGS} from './action_types';

const bookmarkSettings = (state = {}, action: any) => {
    switch (action.type) {
    case RECEIVED_BOOKMARK_SETTINGS: {
        return {
            ...state,
            meeting: action.data,
        };
    }
    default:
        return state;
    }
};

export default combineReducers({
    bookmarkSettings,
});
