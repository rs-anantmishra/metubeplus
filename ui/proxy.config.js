const PROXY_CONFIG = {
    "/youtube.com": {
        "target": "http://localhost:8484",
        "secure": false,
        "changeOrigin": true,
        "logLevel": "debug"
    },
    "/api": {
        "target": "http://localhost:3000",
        "secure": false,
        "changeOrigin": true,
        "logLevel": "debug"
    },
    // "/products": {
    //     "target": "http://localhost:8083",
    //     "secure": false,
    //     "changeOrigin": true,
    //     "logLevel": "debug"
    // },
    // "/business": {
    //     "target": "http://localhost:8081",
    //     "secure": false,
    //     "changeOrigin": true,
    //     "logLevel": "debug",
    // }
};

module.exports = PROXY_CONFIG