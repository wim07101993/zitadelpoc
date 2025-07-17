export const authConfig = {
  clientId: '329309089064615939',
  redirectUri: new URL('http://localhost:8765/'),
  postLogoutRedirectUri: new URL('http://localhost:8765/'),
  authBaseUrl: new URL('http://localhost:7003'),
  getAuthorizationEndpoint: () => new URL(`${authConfig.authBaseUrl}oauth/v2/authorize`),
  getEndSessionEndpoint: () => new URL(`${authConfig.authBaseUrl}oidc/v1/end_session`),
  getTokenEndpoint: () => new URL(`${authConfig.authBaseUrl}oauth/v2/token`),
  getUserInfoEndpoint: () => new URL(`${authConfig.authBaseUrl}oidc/v1/userinfo`),
};
export const defaultScopes = [
  'openid',
  'email',
  'profile',
  'offline_access',
];

export const apiConfig = {
  baseUrl: new URL('http://localhost:8765')
}
