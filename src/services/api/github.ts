import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import chalk from 'chalk';
import open from 'open';

import logger from '@/services/logger';
import {
  AuthMethod, IAPI, IService, postPoll,
} from '@/services/api';

export class GithubAPI implements IAPI, IService {
  readonly name: string = 'github';

  readonly authMethod: AuthMethod = 'token';

  readonly apiUrl: string = 'https://api.github.com';

  readonly url: string = 'https://github.com';

  readonly api: AxiosInstance = axios.create({
    baseURL: this.apiUrl,
    timeout: 1000,
  });

  readonly web = axios.create({
    baseURL: this.url,
    headers: {
      Accept: 'application/json',
    },
    withCredentials: true,
  });

  static create(): GithubAPI {
    return new GithubAPI();
  }

  async auth(): Promise<any> {
    const clientID = '5df58bfe3495a524a0d4';
    const response = await this.web.post('/login/device/code', {
      client_id: clientID,
      scope: 'repo',
    });
    logger.info(`${chalk.bold('Wait')} before continuing`);
    logger.info(
      `${chalk.bold('Copy')} this code: ${chalk.bold.cyan(response.data.user_code)}`,
    );
    open(response.data.verification_uri);
    const pollingResponse = await postPoll(
      {
        interval: response.data.interval,
        timeout: response.data.expires_in,
      },
      this.web,
      '/login/oauth/access_token',
      {
        client_id: clientID,
        device_code: response.data.device_code,
        grant_type: 'urn:ietf:params:oauth:grant-type:device_code',
      },
    );
    const accessToken = pollingResponse.data.access_token;
    return accessToken;
  }

  async get(url: string, config?: AxiosRequestConfig): Promise<AxiosResponse> {
    return this.api.get(url, config);
  }

  async zen(): Promise<string> {
    const response = await this.get('/zen');
    return response.data;
  }

  async user(username: string): Promise<AxiosResponse> {
    const response = await this.get(`/users/${username}`);
    return response;
  }
}

export default GithubAPI;
