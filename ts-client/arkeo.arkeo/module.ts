// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgBondProvider } from "./types/arkeo/arkeo/tx";
import { MsgModProvider } from "./types/arkeo/arkeo/tx";
import { MsgCloseContract } from "./types/arkeo/arkeo/tx";
import { MsgOpenContract } from "./types/arkeo/arkeo/tx";
import { MsgClaimContractIncome } from "./types/arkeo/arkeo/tx";

import { Provider as typeProvider} from "./types"
import { Contract as typeContract} from "./types"
import { ContractExpiration as typeContractExpiration} from "./types"
import { ContractExpirationSet as typeContractExpirationSet} from "./types"
import { ProtoInt64 as typeProtoInt64} from "./types"
import { ProtoUint64 as typeProtoUint64} from "./types"
import { ProtoAccAddresses as typeProtoAccAddresses} from "./types"
import { ProtoStrings as typeProtoStrings} from "./types"
import { ProtoBools as typeProtoBools} from "./types"
import { Params as typeParams} from "./types"

export { MsgBondProvider, MsgModProvider, MsgCloseContract, MsgOpenContract, MsgClaimContractIncome };

type sendMsgBondProviderParams = {
  value: MsgBondProvider,
  fee?: StdFee,
  memo?: string
};

type sendMsgModProviderParams = {
  value: MsgModProvider,
  fee?: StdFee,
  memo?: string
};

type sendMsgCloseContractParams = {
  value: MsgCloseContract,
  fee?: StdFee,
  memo?: string
};

type sendMsgOpenContractParams = {
  value: MsgOpenContract,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimContractIncomeParams = {
  value: MsgClaimContractIncome,
  fee?: StdFee,
  memo?: string
};


type msgBondProviderParams = {
  value: MsgBondProvider,
};

type msgModProviderParams = {
  value: MsgModProvider,
};

type msgCloseContractParams = {
  value: MsgCloseContract,
};

type msgOpenContractParams = {
  value: MsgOpenContract,
};

type msgClaimContractIncomeParams = {
  value: MsgClaimContractIncome,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgBondProvider({ value, fee, memo }: sendMsgBondProviderParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBondProvider: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBondProvider({ value: MsgBondProvider.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBondProvider: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgModProvider({ value, fee, memo }: sendMsgModProviderParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgModProvider: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgModProvider({ value: MsgModProvider.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgModProvider: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCloseContract({ value, fee, memo }: sendMsgCloseContractParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCloseContract: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCloseContract({ value: MsgCloseContract.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCloseContract: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgOpenContract({ value, fee, memo }: sendMsgOpenContractParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgOpenContract: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgOpenContract({ value: MsgOpenContract.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgOpenContract: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaimContractIncome({ value, fee, memo }: sendMsgClaimContractIncomeParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaimContractIncome: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaimContractIncome({ value: MsgClaimContractIncome.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaimContractIncome: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgBondProvider({ value }: msgBondProviderParams): EncodeObject {
			try {
				return { typeUrl: "/arkeo.arkeo.MsgBondProvider", value: MsgBondProvider.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBondProvider: Could not create message: ' + e.message)
			}
		},
		
		msgModProvider({ value }: msgModProviderParams): EncodeObject {
			try {
				return { typeUrl: "/arkeo.arkeo.MsgModProvider", value: MsgModProvider.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgModProvider: Could not create message: ' + e.message)
			}
		},
		
		msgCloseContract({ value }: msgCloseContractParams): EncodeObject {
			try {
				return { typeUrl: "/arkeo.arkeo.MsgCloseContract", value: MsgCloseContract.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCloseContract: Could not create message: ' + e.message)
			}
		},
		
		msgOpenContract({ value }: msgOpenContractParams): EncodeObject {
			try {
				return { typeUrl: "/arkeo.arkeo.MsgOpenContract", value: MsgOpenContract.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgOpenContract: Could not create message: ' + e.message)
			}
		},
		
		msgClaimContractIncome({ value }: msgClaimContractIncomeParams): EncodeObject {
			try {
				return { typeUrl: "/arkeo.arkeo.MsgClaimContractIncome", value: MsgClaimContractIncome.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaimContractIncome: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						Provider: getStructure(typeProvider.fromPartial({})),
						Contract: getStructure(typeContract.fromPartial({})),
						ContractExpiration: getStructure(typeContractExpiration.fromPartial({})),
						ContractExpirationSet: getStructure(typeContractExpirationSet.fromPartial({})),
						ProtoInt64: getStructure(typeProtoInt64.fromPartial({})),
						ProtoUint64: getStructure(typeProtoUint64.fromPartial({})),
						ProtoAccAddresses: getStructure(typeProtoAccAddresses.fromPartial({})),
						ProtoStrings: getStructure(typeProtoStrings.fromPartial({})),
						ProtoBools: getStructure(typeProtoBools.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			ArkeoArkeo: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;