import {CSSProperties, FC, ReactNode} from "react";

export type Overflow = "auto" | "clip" | "hidden" | "scroll" | "visible";

export interface ScrollAreaProps {
    children: ReactNode;
    className?: string;
    style?: CSSProperties;
    overflow?: Overflow;
    overflowX?: Overflow;
    overflowY?: Overflow;
}

const ScrollArea: FC<ScrollAreaProps> = ({children, className, style, overflowY, overflowX, overflow}) => {
    let styles: CSSProperties = {};
    if (overflow) {
        styles = {...styles, ...{overflow}}
    }
    if (overflowX) {
        styles = {...styles, ...{overflowX}}
    }
    if (overflowY) {
        styles = {...styles, ...{overflowY}}
    }
    styles = {...styles, ...style};

    return (
        <div style={styles}
             className={className}>
            {children}
        </div>
    )
}

export default ScrollArea
