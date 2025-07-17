export const authConfig = {
  clientId: '321160767120408579',
  redirectUri: new URL('http://localhost:8765/'),
  authorizationEndpoint: new URL('http://localhost:7003/oauth/v2/authorize'),
  tokenEndpoint: new URL('http://localhost:7003/oauth/v2/token')
};
export const defaultScopes = ['openid', 'email', 'profile', 'offline_access'];

export const apiConfig = {
  baseUrl: new URL('http://localhost:8765')
}