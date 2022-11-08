import { FC, ReactNode, useCallback, useEffect, useRef, useState } from "react";

export interface ResizableContainerProps {
  children: ReactNode;
  width?: number;
  conatinerClassName?: string;
  resizerClassName?: string;
}

const ResizableContainer: FC<ResizableContainerProps> = ({
  children,
  width = 256,
  conatinerClassName = "",
  resizerClassName = "bg-gray-100 hover:bg-gray-200 w-[2px]",
}) => {
  const containerRef = useRef(null);
  const [isResizing, setIsResizing] = useState(false);
  const [sidebarWidth, setSidebarWidth] = useState(width);

  const startResizing = useCallback((mouseDownEvent: any) => {
    setIsResizing(true);
  }, []);

  const stopResizing = useCallback(() => {
    setIsResizing(false);
  }, []);

  const resize = useCallback(
    (mouseMoveEvent: any) => {
      if (isResizing) {
        setSidebarWidth(
          mouseMoveEvent.clientX -
            containerRef.current.getBoundingClientRect().left
        );
      }
    },
    [isResizing]
  );

  useEffect(() => {
    window.addEventListener("mousemove", resize);
    window.addEventListener("mouseup", stopResizing);
    return () => {
      window.removeEventListener("mousemove", resize);
      window.removeEventListener("mouseup", stopResizing);
    };
  }, [resize, stopResizing]);

  return (
    <div
      ref={containerRef}
      className={`${conatinerClassName} flex`}
      style={{ width: sidebarWidth }}
      onMouseDown={(e) => e.preventDefault()}
    >
      <div className="flex-1">{children}</div>
      <div
        className={`${resizerClassName} __resizer`}
        onMouseDown={startResizing}
      />
    </div>
  );
};

export default ResizableContainer;
