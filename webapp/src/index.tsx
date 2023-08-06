import React from 'react';
import {Store, Action} from 'redux';

import {GlobalState} from 'mattermost-redux/types/store';

import reducer from 'reducers';

import manifest from './manifest';
import Client from './client';

// eslint-disable-next-line import/no-unresolved
import {PluginRegistry} from './types/mattermost-webapp';
import {getServerRoute} from './selectors';

import NotepadIcon from './components/icons/icons';
import RHSView from './components/right_hand_sidebar';

export default class Plugin {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-empty-function
    public async initialize(
        registry: PluginRegistry,
        store: Store<GlobalState, Action<Record<string, unknown>>>,
    ) {
        Client.setServerRoute(getServerRoute(store.getState()));

        // @see https://developers.mattermost.com/extend/plugins/webapp/reference/
        registry.registerReducer(reducer);

        // store.dispatch()

        Client.getConfiguration().then((result) => {
            // setResponse(result);
            // eslint-disable-next-line no-console
            console.log(result);
        });

        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(
            RHSView,
            'Notepad',
        );

        registry.registerChannelHeaderButtonAction(
            <NotepadIcon/>,
            () => store.dispatch(toggleRHSPlugin),
            'Notepad',
            'View Notepad',
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
