export class Api {
    /**
     * @param baseUrl {URL}
     */
    constructor(baseUrl) {
        this.baseUrl = baseUrl;
    }

    /**
     * @return {Promise<String>}
     */
    async getSecret() {
        const url = `${this.baseUrl}/api/secret`;

        /**
         * @type Response
         */
        const response = await fetch(url);
        if (response.status >= 500) {
            throw `failed to fetch secret (server error): ${response.status} ${response.statusText}: ${await response.text()}`;
        } else if (response.status >= 400) {
            throw `failed to fetch secret: ${response.status} ${response.statusText}: ${await response.text()}`;
        }

        return await response.text();
    }
}
