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

    const handleModifyClick = useCallback((event: React.MouseEvent<HTMLSpanElement> | undefined): void => {
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
        <div className='row remove-margin'>
            { showEditButton && (
                <div className='pull-right'>
                    <span
                        className='glyphicon glyphicon-pencil'
                        aria-hidden='true'
                        data-toggle='tooltip'
                        data-placement='left'
                        title='수정'
                        onClick={handleModifyClick}
                    />
                </div>
            ) }
            {editMode && (
                <>
                    <TextareaAutosize
                        className='form-control'
                        style={{
                            width: '100%',
                        }}
                        placeholder='여기에 입력하세요.'
                        value={value}
                        onChange={handleValueChange}
                    />
                    <hr/>
                    <div className='row text-center remove-margin'>
                        <button
                            type='button'
                            className='btn btn-secondary'
                            onClick={handleCancelClick}
                        >{'취소'}</button>
                        <button
                            type='button'
                            className='btn btn-primary'
                            onClick={handleSaveClick}
                        >{'저장'}</button>

                    </div>
                </>
            )}
            {!editMode && (
                <>
                    {PostUtils.messageHtmlToComponent(PostUtils.formatText(value))}
                </>
            )}
        </div>
    );
};

export default BodyComponents;
