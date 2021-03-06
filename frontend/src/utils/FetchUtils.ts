export const fetchPublic = async (url: string, requestOptions: RequestInit): Promise<any> => {
    const response = await fetch(url, requestOptions)
    return response.json()
}

export const fetchAuthed = async (url: string, requestOptions: RequestInit) : Promise<any> => {
    requestOptions.headers = {...requestOptions.headers, ...{'Token': getCookie("studyhub_token")}}
    const response = await fetch(url, requestOptions)
    return response.json()
}

const getCookie = (cookieName: string): string  => {
    let name = cookieName + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i < ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) === ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) === 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}