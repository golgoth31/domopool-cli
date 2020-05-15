const path = require('path');
const ProgressBarPlugin = require('progress-bar-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const DashboardPlugin = require('webpack-dashboard/plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');

module.exports = {
  mode: "production",
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'js/bundle.js',
    // filename: '[name].bundle.js',
    // chunkFilename: '[name].bundle.js',
  },
  // Enable sourcemaps for debugging webpack's output.
  // devtool: "source-map",

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

  optimization: {
    // splitChunks: {
    //   chunks: 'all',
    //   maxSize: 100000
    // },
    // runtimeChunk: {
    //   name: 'runtime'
    // },
    minimize: true,
    minimizer: [
      new TerserPlugin({
        parallel: true,
        sourceMap: false,
        terserOptions: {
          output: {
            comments: false
          }
        }
      }),
      new OptimizeCSSAssetsPlugin({})
    ]
  },
  plugins: [
    new ProgressBarPlugin({ incomplete: '-' }),
    new CleanWebpackPlugin(),
    new DashboardPlugin(),
    new HtmlWebpackPlugin({
      template: 'src/index.html'
    }),
  ],
};
