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
        <div className='hsan-header' >
            <div className='header-left'>
                <select onChange={handleSelectBoxChange}>
                    <option value='channel'>{'채널'}</option>
                    <option value='common'>{'공통'}</option>
                </select>
            </div>
            <div className='header-right'>
                <button
                    type='button'
                    onClick={handleReloadClick}
                >{'새로고침'}</button>
            </div>
        </div>
    );
};

export default HeaderComponents;
