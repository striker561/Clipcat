
export const insertLinks = (text: string | undefined) => {
    if (!text) return text;
    const urlRegex = /^(?:([A-Za-z]+):)?(\/{0,3})([0-9.\-A-Za-z]+)(?::(\d+))?(?:\/([^?#]*))?(?:\?([^#]*))?(?:#(.*))?$/;

    const matches = text.match(urlRegex);

    let modifiedText = text;

    if (matches) {
        if (matches.length > 0) {
            for (let i = 0; i < matches.length; i++) {
                const url = matches[i];
                const link = `<a href="${url}" target="_blank" rel="noopener noreferrer">${url}</a>`;
                modifiedText = modifiedText.replace(url, link);
            }
        }
    }

    return modifiedText;
}