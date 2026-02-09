
export const insertLinks = (text: string | any) => {
    if (!text) return text;
    const urlRegex = /https?:\/\/(?:www\.)?[^\s/$.?#].[^\s]*/;


    const matches = text.match(urlRegex);

    let modifiedText = text;

    if (matches) {
        if (matches.length > 0) {
            for (let i = 0; i < matches.length; i++) {
                const url = matches[i];
                const link = `<a href="${url}" target="_blank" class="inserted-link">${url}</a>`;
                modifiedText = modifiedText.replace(url, link);
            }
        }
    }

    return modifiedText;
}