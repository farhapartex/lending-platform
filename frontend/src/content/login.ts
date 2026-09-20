export const loginContent = {
  eyebrow: "Connect",
  title: "Your wallet is your account.",
  description:
    "There is no sign up, no password and no email. Connecting a wallet proves an address belongs to you, and that is the whole of it.",
  points: [
    {
      title: "Nothing is created",
      body: "No account record, no profile, no credentials to lose. Disconnect and there is nothing left behind on our side.",
    },
    {
      title: "We cannot move your funds",
      body: "Connecting grants permission to read your address. Every transfer needs a separate signature that you approve in your wallet.",
    },
    {
      title: "We never see your keys",
      body: "Your recovery phrase and private keys stay in your wallet. They are never sent to this site and we could not use them if they were.",
    },
  ],
  panelTitle: "Choose a wallet",
  panelDescription: "MetaMask is the only option for now. More will follow.",
  footnote: "By connecting you agree that this is unaudited software and not somewhere to put money you cannot lose.",
  redirectingLabel: "Wallet connected, taking you in",
  notDetectedTitle: "MetaMask was not detected",
  notDetectedBody:
    "Install the MetaMask extension, or open this page inside a wallet browser, then reload and try again.",
  installLabel: "Get MetaMask",
  installUrl: "https://metamask.io/download/",
} as const;
