import {authorize, getUserInfo} from "./auth.js";
import {Api} from "./api.js";
import {apiConfig, authConfig} from "./config.js";

const api = new Api(apiConfig.baseUrl);

/**
 * @type {String | null}
 */
let accessToken = null

async function updateContent() {
  console.log(await getUserInfo(authConfig.userInfoEndpoint));
  const contentElement = document.getElementById('content');
  if (contentElement == null) {
    return;
  }

  if (accessToken == null) {
    const button = document.createElement('button');
    button.addEventListener('click', login);
    button.innerHTML = 'Log in';
    contentElement.innerHTML = '<p>You are not yet authenticated. please authenticate to see the secret.</p>';
    contentElement.appendChild(button);
    return;
  }

  const secret = await api.getSecret(accessToken);
  contentElement.innerHTML = `
        <p>Congratulations. Here is your secret.</p>
        <p>${secret}</p>
    `;
}

export async function login() {
  accessToken = await authorize(true);
  await updateContent();
}

async function main() {
  accessToken = await authorize(false);
  await updateContent();
}

await main();
