/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable no-console */
import React, {useCallback, useEffect, useState} from 'react';

import {useSelector} from 'react-redux';

import {GlobalState} from 'mattermost-redux/types/store';
import {Channel} from 'mattermost-redux/types/channels';
import {getCurrentChannelId} from 'mattermost-redux/selectors/entities/channels';

// eslint-disable-next-line import/no-unresolved
import {Response} from 'types';

import Client from 'client';

import './style.scss';

import HeaderComponents from './header/HeaderComponents';
import BodyComponents from './body/BodyComponents';

const RHSView = () => {
    const [response, setResponse] = useState<Response | null>(null);
    const [selectBox, setSelectBox] = useState<string>('channel');
    const [reload, toggleReload] = useState<boolean>(true);

    const currentChannelId = useSelector<GlobalState, string>((state) => getCurrentChannelId(state));

    useEffect(() => {
        if (currentChannelId) {
            Client.getBookmark(currentChannelId).then((result) => {
                setResponse(result);
            });
        }
    }, [currentChannelId, reload]);

    const handleSelectBoxChange = useCallback((event: React.ChangeEvent<HTMLSelectElement>): void => {
        const {value} = event.target;
        setSelectBox(value);
    }, []);

    const handleReloadClick = useCallback((event: React.MouseEvent<HTMLButtonElement> | undefined): void => {
        toggleReload(!reload);
    }, [reload]);

    return (
        <div
            className='hsan-right-hand-sidebar-wrapper'
        >
            {response && (
                <div>
                    <HeaderComponents
                        handleSelectBoxChange={handleSelectBoxChange}
                        handleReloadClick={handleReloadClick}
                    />
                    <hr/>
                    <BodyComponents
                        selectBox={selectBox}
                        response={response}
                    />

                </div>
            )}
        </div>
    );
};

export default RHSView;
