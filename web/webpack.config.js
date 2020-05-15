const path = require('path');
const ProgressBarPlugin = require('progress-bar-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const DashboardPlugin = require('webpack-dashboard/plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');

module.exports = {
  mode: "production",
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'js/bundle.js'
  },
  // Enable sourcemaps for debugging webpack's output.
  devtool: "source-map",

  resolve: {
    // Add '.ts' and '.tsx' as resolvable extensions.
    extensions: [".ts", ".tsx", '.js']
  },

  devServer: {
    contentBase: path.join(__dirname, 'build'),
    compress: true,
    port: 8000,
    proxy: {
      '/': {
        target: 'http://localhost:8080',
        secure: false,
        changeOrigin: true
      }
    }
  },

  module: {
    rules: [
      {
        test: /\.ts(x?)$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "ts-loader"
          }
        ]
      },
      // All output '.js' files will have any sourcemaps re-processed by 'source-map-loader'.
      {
        enforce: "pre",
        test: /\.js$/,
        loader: "source-map-loader"
      },
      {
        test: /\.html$/,
        loader: 'html-loader'
      }
    ]
  },

  // When importing a module whose path matches one of the following, just
  // assume a corresponding global variable exists and use that instead.
  // This is important because it allows us to avoid bundling all of our
  // dependencies, which allows browsers to cache those libraries between builds.
  // externals: {
  //   react: "React",
  //   "react-dom": "ReactDOM"
  // },
  plugins: [
    new ProgressBarPlugin({ incomplete: '-' }),
    new CleanWebpackPlugin(),
    new DashboardPlugin(),
    new HtmlWebpackPlugin({
      template: 'src/index.html'
    }),
  ],
};
