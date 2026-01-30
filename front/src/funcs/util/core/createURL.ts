function createURL(url: string): string {
  const baseUrl = import.meta.env.VITE_API_URL;
  const separator = baseUrl.endsWith('/') ? '' : '/';
  return baseUrl + separator + url;
}

export default createURL;
