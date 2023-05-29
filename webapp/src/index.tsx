import React from 'react';
import {Store, Action} from 'redux';

import {GlobalState} from 'mattermost-redux/types/store';

import manifest from './manifest';
import Client from './client';

// eslint-disable-next-line import/no-unresolved
import {PluginRegistry} from './types/mattermost-webapp';
import {getServerRoute} from './selectors';

import reducer from './reducer';

import BookmarkIcon from './components/icons/icons';
import RHSView from './components/right_hand_sidebar';

export default class Plugin {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-empty-function
    public async initialize(
        registry: PluginRegistry,
        store: Store<GlobalState, Action<Record<string, unknown>>>,
    ) {
        Client.setServerRoute(getServerRoute(store.getState()));

        // @see https://developers.mattermost.com/extend/plugins/webapp/reference/
        // registry.registerReducer(reducer);

        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(
            RHSView,
            'Bookmark',
        );

        registry.registerChannelHeaderButtonAction(
            <BookmarkIcon/>,
            () => store.dispatch(toggleRHSPlugin),
            'Bookmark',
            'View Bookmark',
        );
    }

    uninitialize() {
    // No clean up required.
    }
}

declare global {
    interface Window {
        registerPlugin(id: string, plugin: Plugin): void;
    }
}

window.registerPlugin(manifest.id, new Plugin());
