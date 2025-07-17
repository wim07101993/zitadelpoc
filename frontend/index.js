import {authorize, endSession, getUserInfo} from "./auth.js";
import {Api} from "./api.js";
import {apiConfig, authConfig} from "./config.js";

const api = new Api(apiConfig.baseUrl);

/**
 * @type {String | null}
 */
let accessToken = null

function loadLoginPage() {
  const loginButton = document.createElement('button');
  loginButton.addEventListener('click', login);
  loginButton.innerHTML = 'Log in';

  const contentElement = document.getElementById('content');
  contentElement.innerHTML = `
    <p>You are not yet authenticated. please authenticate to see the secret.</p>
  `;
  contentElement.appendChild(loginButton);
}

function loadSecretPage(secret) {
  const logoutButton = document.createElement('button');
  logoutButton.addEventListener('click', logout);
  logoutButton.innerHTML = 'Log out';

  const contentElement = document.getElementById('content');
  contentElement.innerHTML = `
    <p>Congratulations. Here is your secret.</p>
    <p>${secret}</p>
  `;
  contentElement.appendChild(logoutButton);
}

async function updateContent() {
  if (accessToken == null) {
    loadLoginPage();
    return;
  }

  const secret = await api.getSecret(accessToken);
  loadSecretPage(secret);
}

export async function login() {
  accessToken = await authorize(true);
  await updateContent();
}

async function logout(){
  await endSession();
  await updateContent();
}

async function main() {
  accessToken = await authorize(false);
  await updateContent();
}

await main();
