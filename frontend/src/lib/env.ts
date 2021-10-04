const env = import.meta.env.VITE_ENV as string

const configProd = {
    baseWS: "wss://matrix.sprinteins.com/poker/api",
}

type Config = typeof configProd;

const configDev: Config = {
    ...configProd,
    baseWS: "ws://localhost:7788/poker/api",
}

const configMap: {[key: string]: Config} = {
    "PROD": configProd,
    "DEV":  configDev,
}

export let config = configMap[env];
if(!config){
    config = configProd
}



