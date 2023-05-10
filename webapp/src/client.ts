/* eslint-disable import/no-named-as-default */
/* eslint-disable import/no-unresolved */
import {Client4} from 'mattermost-redux/client';
import {ClientError} from 'mattermost-redux/client/client4';

import {id as pluginId} from './manifest';

// export default class Client {
//     constructor() {
//         this.url = `/plugins/${pluginId}`;
//     }

//     getBookmarkSettings = async () => {
//         return this.doGet(`${this.url}/bookmark`);
//     };

//     // saveMeetingSettings = async (meeting) => {
//     //     return this.doPost(`${this.url}/settings`, meeting);
//     // };

//     doGet = async (url, headers = {}) => {
//         return this.doFetch(url, {headers});
//     };

//     doPost = async (url, body, headers = {}) => {
//         return this.doFetch(url, {
//             method: 'POST',
//             body: JSON.stringify(body),
//             headers: {
//                 ...headers,
//                 'Content-Type': 'application/json',
//             },
//         });
//     };

//     doFetch = async (url, {method = 'GET', body = null, headers = {}}) => {
//         const options = Client4.getOptions({
//             method,
//             body,
//             headers: {
//                 ...headers,
//                 Accept: 'application/json',
//             },
//         });

//         const response = await fetch(url, options);

//         if (response.ok) {
//             return response.json();
//         }

//         const data = await response.text();

//         throw new ClientError(Client4.url, {
//             message: data || '',
//             status_code: response.status,
//             url,
//         });
//     };
// }

class ClientClass {
    url = '';

    setServerRoute(url: string) {
        this.url = url + `/plugins/${pluginId}`;
    }

    needsConnect = async () => {
        return this.doGet(`${this.url}/needsConnect`);
    };

    getBookmark = async () => {
        return this.doGet(`${this.url}/bookmark`);
    };

    doGet = async (url: string, headers: {[key: string]: any} = {}) => {
        headers['X-Timezone-Offset'] = new Date().getTimezoneOffset();

        const options = {
            method: 'get',
            headers,
        };

        const response = await fetch(url, Client4.getOptions(options));

        if (response.ok) {
            return response.json();
        }

        const text = await response.text();

        throw new ClientError(Client4.url, {
            message: text || '',
            status_code: response.status,
            url,
        });
    };

    doPost = async (url: string, body: any = {}, headers: {[key: string]: any} = {}) => {
        headers['X-Timezone-Offset'] = new Date().getTimezoneOffset();

        const options = {
            method: 'post',
            body: JSON.stringify(body),
            headers,
        };

        const response = await fetch(url, Client4.getOptions(options));

        if (response.ok) {
            return response.json();
        }

        const text = await response.text();

        throw new ClientError(Client4.url, {
            message: text || '',
            status_code: response.status,
            url,
        });
    };

    doDelete = async (url: string, headers: {[key: string]: any} = {}) => {
        headers['X-Timezone-Offset'] = new Date().getTimezoneOffset();

        const options = {
            method: 'delete',
            headers,
        };

        const response = await fetch(url, Client4.getOptions(options));

        if (response.ok) {
            return response.json();
        }

        const text = await response.text();

        throw new ClientError(Client4.url, {
            message: text || '',
            status_code: response.status,
            url,
        });
    };

    doPut = async (url: string, body: any, headers: {[key: string]: any} = {}) => {
        headers['X-Timezone-Offset'] = new Date().getTimezoneOffset();

        const options = {
            method: 'put',
            body: JSON.stringify(body),
            headers,
        };

        const response = await fetch(url, Client4.getOptions(options));

        if (response.ok) {
            return response.json();
        }

        const text = await response.text();

        throw new ClientError(Client4.url, {
            message: text || '',
            status_code: response.status,
            url,
        });
    };
}

const Client = new ClientClass();

export default Client;
