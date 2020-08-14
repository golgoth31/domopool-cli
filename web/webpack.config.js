const path = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const DashboardPlugin = require('webpack-dashboard/plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const FaviconsWebpackPlugin = require('favicons-webpack-plugin');
const HardSourceWebpackPlugin = require('hard-source-webpack-plugin');

module.exports = {
  mode: "production",
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'bundle.js',
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
        target: 'http://192.168.11.183',
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
        include: path.resolve(__dirname, 'src'),
        use: [
          {
            loader: "ts-loader"
          }
        ]
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        include: path.resolve(__dirname, 'src'),
        loader: 'babel-loader'
      },
      // All output '.js' files will have any sourcemaps re-processed by 'source-map-loader'.
      {
        enforce: "pre",
        test: /\.js$/,
        include: path.resolve(__dirname, 'src'),
        loader: "source-map-loader"
      },
      {
        test: /\.html$/,
        include: path.resolve(__dirname, 'src'),
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
    new CleanWebpackPlugin(),
    new DashboardPlugin(),
    new HtmlWebpackPlugin({
      template: 'src/index.html'
    }),
    new FaviconsWebpackPlugin({
      logo: './src/favicon.png',
      prefix: '/',
    }),
    new HardSourceWebpackPlugin()
  ],
};
