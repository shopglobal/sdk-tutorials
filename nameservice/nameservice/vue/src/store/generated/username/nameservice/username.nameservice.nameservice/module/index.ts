// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateWhois } from "./types/nameservice/tx";
import { MsgDeleteWhois } from "./types/nameservice/tx";
import { MsgCreateWhois } from "./types/nameservice/tx";


const types = [
  ["/username.nameservice.nameservice.MsgUpdateWhois", MsgUpdateWhois],
  ["/username.nameservice.nameservice.MsgDeleteWhois", MsgDeleteWhois],
  ["/username.nameservice.nameservice.MsgCreateWhois", MsgCreateWhois],
  
];

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgUpdateWhois: (data: MsgUpdateWhois): EncodeObject => ({ typeUrl: "/username.nameservice.nameservice.MsgUpdateWhois", value: data }),
    msgDeleteWhois: (data: MsgDeleteWhois): EncodeObject => ({ typeUrl: "/username.nameservice.nameservice.MsgDeleteWhois", value: data }),
    msgCreateWhois: (data: MsgCreateWhois): EncodeObject => ({ typeUrl: "/username.nameservice.nameservice.MsgCreateWhois", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};