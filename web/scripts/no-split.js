const rewire = require('rewire');
const defaults = rewire('react-scripts/scripts/build.js');
let config = defaults.__get__('config');

config.optimization.splitChunks = {
    cacheGroups: {
        default: false,
    },
};

config.optimization.runtimeChunk = false;
config.output.filename = 'bundle.js';
config.plugins[5].options.filename = 'bundle.css';
config.plugins[5].options.moduleFilename = () => 'bundle.css';
