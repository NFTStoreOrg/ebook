import { defineStore } from 'pinia'
import MetaMaskSDK, { SDKProvider } from '@metamask/sdk';

export const useSignatureStore = defineStore('personal_sign', {
  //函式
  actions: {
    async onConnect() {
      try {
        await this.sdk?.connect();
        this.lastResponse = '';
        this.provider = this.sdk?.getProvider();
        this.chainId = this.provider?.getChainId();
        this.account = this.provider?.getSelectedAddress();
        //  Chain changed
        this.provider?.on('chainChanged', chain => {
          console.log(`App::Chain changed:'`, chain);
          this.chainId = chain;
        });

        //  Accounts changed
        this.provider?.on('accountsChanged', accounts => {
          let accoontArray = accounts as string[];
          console.log(`App::Accounts changed:'`, accounts);
          this.account = accoontArray[0]
        })

        //  Connected event
        this.provider?.on('connect', _connectInfo => {
          console.log(`App::connect`, _connectInfo)
          this.onConnect()
          this.connected = true
        })

        //  Disconnected event
        this.provider?.on('disconnect', error => {
          console.log(`App::disconnect`, error)
          this.connected = false
        })

        this.availableLanguages = this.sdk?.availableLanguages ?? ['en'];
      } catch (err) {
        console.log(`erquest accounts error`, err)
      }
    },
    async addPolygonChain() {
      try {
        const res = await this.provider?.request({
          method: 'wallet_addEthereumChain',
          params: [
            {
              chainId: '0x89',
              chainName: 'Polygon',
              blockExplorerUrls: ['https://polygonscan.com'],
              nativeCurrency: { symbol: 'MATIC', decimals: 18 },
              rpcUrls: ['https://polygon-rpc.com/'],
            },
          ],
        })
        console.log('add', res)
        this.lastResponse = res as string
      } catch (err) {
        console.log('Add error', err)
      }
    },
    async getSign() {
      if (!this.provider) {
        this.onConnect()
      }
      if (!this.connected) {
        this.onConnect()
      }
      try {
        const from = this.provider?.getSelectedAddress();
        const message = `Welcome to YiSin ebook store!

Click to verify that you own this wallet and have control over it.

YiSin ebook (https://yisinnft.org/ebook) need to confirm whether you have the permission to read the e-book file.

This request will not trigger a blockchain transaction or cost any gas fees.`;
        const hexMessage = '0x' + Buffer.from(message, 'utf8').toString('hex');

        const sign = await window.ethereum?.request({
          method: 'personal_sign',
          params: [hexMessage, from, 'Example password'],
        });
        console.log(`sign: ${sign}`);
        return sign;
      } catch (err: any) {
        console.log(err);
        return 'Error: ' + err.message;
      }
    },
  },
  //數據儲存
  state() {
    var sdk = new MetaMaskSDK({
      dappMetadata: {
        url: window.location.href,
        name: 'Yi Sin ebook',
      },
      enableAnalytics: true,
      checkInstallationImmediately: false,
      logging: {
        developerMode: true,
        sdk: true
      },
      i18nOptions: {
        enabled: true,
      }
    })
    var provider = sdk.getProvider() as SDKProvider | undefined
    var availableLanguages: string[] = ['']
    return {
      sdk,
      account: provider?.getSelectedAddress() as string | null | undefined,
      chainId: provider?.getChainId() as string | undefined | unknown,
      connected: false,
      lastResponse: '',
      provider,
      availableLanguages,
      selectedLanguage: '',
    }
  }
})