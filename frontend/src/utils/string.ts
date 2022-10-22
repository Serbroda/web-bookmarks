const hashString = (str: string) => {
    let hash = 0,
        i,
        chr;
    if (!str || str.length === 0) {
        return hash;
    }
    for (i = 0; i < str.length; i++) {
        chr = str.charCodeAt(i);
        hash = (hash << 5) - hash + chr;
        hash |= 0; // Convert to 32bit integer
    }
    return hash;
};

const getSecondPart = (str: string, character: string): string | undefined => {
    const parts = str.split(character);
    if (parts && parts.length > 1) {
        return parts[1];
    }
    return undefined;
};

const getFirstPart = (str: string, character: string) => {
    const parts = str.split(character);
    if (parts && parts.length > 0) {
        return parts[0];
    }
    return undefined;
};

export { hashString, getSecondPart, getFirstPart };
