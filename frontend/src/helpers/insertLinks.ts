
export const insertLinks = (text: string | any) => {
    if (!text) return text;
    
    // First, escape all HTML entities to prevent arbitrary HTML injection
    const escapeHtml = (str: string) => {
        return str
            .replace(/&/g, '&amp;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#039;');
    };
    
    // Escape the original text
    let modifiedText = escapeHtml(text);
    
    // Now find URLs and convert them to links
    const urlRegex = /https?:\/\/(?:www\.)?[^\s/$.?#].[^\s]*/g;
    const matches = modifiedText.match(urlRegex);

    if (matches) {
        for (let i = 0; i < matches.length; i++) {
            const url = matches[i];
            const link = `<a href="${url}" target="_blank" class="inserted-link">${url}</a>`;
            modifiedText = modifiedText.replace(url, link);
        }
    }

    return modifiedText;
}