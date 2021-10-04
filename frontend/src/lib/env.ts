const env = import.meta.env.VITE_ENV as string

const configProd = {
    baseWS: "wss://matrix.sprinteins.com/poker",
}

type Config = typeof configProd;

const configDev: Config = {
    ...configProd,
    baseWS: "ws://localhost:7788",
}

const configMap: {[key: string]: Config} = {
    "PROD": configProd,
    "DEV":  configDev,
}

export let config = configMap[env];
if(!config){
    config = configProd
}

console.debug('[DEBUG] ', {env} )


