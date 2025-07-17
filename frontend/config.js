export const authConfig = {
  clientId: '329309089064615939',
  redirectUri: new URL('http://localhost:8765/'),
  authorizationEndpoint: new URL('http://localhost:7003/oauth/v2/authorize'),
  tokenEndpoint: new URL('http://localhost:7003/oauth/v2/token'),
  userInfoEndpoint: new URL('http://localhost:7003/oidc/v1/userinfo')
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
