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

export { hostname, faviconUrl };
