import * as React from 'react';

export default class Icon extends React.PureComponent {
    render() {
        const style = getStyle();
        return (
            <span
                style={style.iconStyle}
                className='icon'
                aria-hidden='true'
            />
        );
    }
}

function getStyle(): { [key: string]: React.CSSProperties } {
    return {
        iconStyle: {
            position: 'relative',
            top: '-1px'
        }
    };
}
