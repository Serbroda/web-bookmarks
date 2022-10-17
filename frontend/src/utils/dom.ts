const classNames = (...classes: string[]): string => {
    return classes.filter(Boolean).join(" ");
};

export { classNames };
