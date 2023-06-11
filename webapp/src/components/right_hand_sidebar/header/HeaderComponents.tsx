/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable no-console */
import React, {useCallback, useState} from 'react';

interface Props {
    handleSelectBoxChange: (event: React.ChangeEvent<HTMLSelectElement>) => void;

    handleReloadClick: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

const HeaderComponents = ({
    handleSelectBoxChange,
    handleReloadClick,
}:Props) => {
    console.log('hsan header components rednering');

    return (
        <div className='row remove-margin'>
            <div className='col-xs-8'>
                <select
                    className='form-control'
                    onChange={handleSelectBoxChange}
                >
                    <option value='channel'>{'채널'}</option>
                    <option value='common'>{'공통'}</option>
                </select>
            </div>
            <div className='col-xs-4'>
                <button
                    type='button'
                    className='btn btn-default'
                    onClick={handleReloadClick}
                >
                    <span
                        className='glyphicon glyphicon-refresh'
                        aria-hidden='true'
                    />
                    {'새로고침'}</button>

            </div>
        </div>
    );
};

export default HeaderComponents;
