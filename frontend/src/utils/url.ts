import { getSecondPart, getFirstPart } from "./string";

const hostname = (url: string): string | undefined => {
    try {
        const u = new URL(url);

        let hostname = u.hostname;
        hostname = hostname.replace(/\/\s*$/, "");

        return hostname;
    } catch (err) {
        return undefined;
    }
};

const faviconUrl = (hostname: string | undefined): string | undefined => {
    if (!hostname) {
        return undefined;
    }
    return `https://icons.duckduckgo.com/ip3/${hostname}.ico`;
};

const removeQuery = (url: string, param: string): string | undefined => {
    const first = getFirstPart(url, "?");
    const second = getSecondPart(url, "?");

    if (second) {
        const ps = new URLSearchParams(second);
        ps.delete(param);

        const search = ps.toString();
        if (search) {
            return `${first}?${search}`;
        } else {
            return `${first}`;
        }
    } else {
        return `${first}`;
    }
};

const addQuery = (url: string, param: string, value: string): string => {
    const first = getFirstPart(url, "?");
    const second = getSecondPart(url, "?");

    const ps = new URLSearchParams(second);
    if (ps.get(param)) {
        ps.set(param, value);
    } else {
        ps.append(param, value);
    }

    const search = ps.toString();
    if (search) {
        return `${first}?${search}`;
    } else {
        return `${first}`;
    }
};

const pushHistoryState = (url: string): void => {
    window.history.pushState({}, document.title, url);
};

export { hostname, faviconUrl, removeQuery, addQuery, pushHistoryState };
