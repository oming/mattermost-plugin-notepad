/* eslint-disable no-console */
import React, {useEffect, useCallback, useState} from 'react';
import TextareaAutosize from 'react-textarea-autosize';
// eslint-disable-next-line import/no-unresolved
import {Request, Response} from 'types';

import Client from 'client';

// @ts-ignore
const PostUtils = window.PostUtils; // must be accessed through `window`

interface Props {
    response: Response;
    selectBox: string;
}

const BodyComponents = ({response, selectBox}: Props) => {
    console.log('hsan body components rednering', response, selectBox);
    const [value, setValue] = useState<string>(response.channelNotepad);
    const [showEditButton, setShowEditButton] = useState<boolean>(selectBox === 'channel');
    const [editMode, setEditMode] = useState<boolean>(false);

    useEffect(() => {
        if (selectBox) {
            if (selectBox === 'channel') {
                setValue(response.channelNotepad);
                setShowEditButton(true);
            } else {
                setValue(response.commonNotepad);
                setShowEditButton(false);
            }
            setEditMode(false);
        }
    }, [response.channelNotepad, response.commonNotepad, selectBox]);

    const handleValueChange = useCallback((event: React.ChangeEvent<HTMLTextAreaElement>): void => {
        if (event) {
            setValue(event.target.value);
        }
    }, []);

    const handleModifyClick = useCallback((event: React.MouseEvent<HTMLButtonElement> | undefined): void => {
        if (event) {
            setEditMode(true);
        }
    }, []);

    const handleCancelClick = useCallback((event: React.MouseEvent<HTMLButtonElement> | undefined): void => {
        if (event) {
            setEditMode(false);
            setValue(response.channelNotepad);
        }
    }, [response.channelNotepad]);

    const handleSaveClick = useCallback((event: React.MouseEvent<HTMLButtonElement> | undefined): void => {
        if (event) {
            setEditMode(false);
            const payload: Request = {
                channelId: response.channelId,
                notepad_content: value,
            };

            Client.saveNotepad(payload).then((result) => {
                console.log('저장함: ', result);
            });
        }
    }, [response.channelId, value]);

    return (
        <div className='hsan-body'>
            { showEditButton && (
                <button
                    type='button'
                    onClick={handleModifyClick}
                >{'수정'}</button>
            ) }
            {editMode && (
                <div
                    className='body-contents'
                >
                    <TextareaAutosize
                        style={{
                            width: '100%',
                        }}
                        placeholder='여기에 입력하세요.'
                        value={value}
                        onChange={handleValueChange}
                    />
                    <hr/>
                    <button
                        type='button'
                        onClick={handleCancelClick}
                    >{'취소'}</button>
                    <button
                        type='button'
                        onClick={handleSaveClick}
                    >{'저장'}</button>
                </div>
            )}
            {!editMode && (
                <div className='body-contents'>
                    {PostUtils.messageHtmlToComponent(PostUtils.formatText(value))}
                </div>
            )}
        </div>
    );
};

export default BodyComponents;
