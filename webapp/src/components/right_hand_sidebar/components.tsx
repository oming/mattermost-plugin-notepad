/* eslint-disable no-console */
import React, {useEffect, useState} from 'react';

import Client from 'client';

// import {useIntl} from 'react-intl';

// @ts-ignore
const PostUtils = window.PostUtils; // must be accessed through `window`

interface Props {
    contents: string // BookmarkContent
    fetchBookmarkSettings: () => void,
}

// const RHSView = ({contents, fetchBookmarkSettings}: Props) => {
const RHSView = () => {
    const [contents, setContents] = useState('');

    useEffect(() => {
        console.log('컴포넌트가 화면에 나타남');
        Client.getBookmark().then((result) => {
            console.log('hsan', result);
            setContents(result.bookmark);
        });
        return () => {
            console.log('컴포넌트가 화면에서 사라짐');
        };
    }, []);

    // const {formatMessage} = useIntl();

    return (
        <div
            style={{
                padding: '10px',
                overflow: 'scroll',
            }}
        >
            {PostUtils.messageHtmlToComponent(PostUtils.formatText(contents))}
        </div>
    );
};

export default RHSView;
