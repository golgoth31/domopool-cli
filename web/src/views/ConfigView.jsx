import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
    Formik,
    Form,
} from 'formik';
import NetworkForm from '../forms/NetworkForm';
import GlobalForm from '../forms/GlobalForm';
import {
    LinearProgress,
    Card,
    CardContent,
    Button
} from "@mui/material";
import TempForm from '../forms/TempForm';

var domopool_pb = require('../proto/domopool_pb');

export default class ConfigView extends React.Component {
    constructor(
        conf
    ) {
        super(conf);
        this.state = {
            config: this.conf,
            ipDisabled: false,
        }
    };

    componentDidMount() {
        dataprovider.get(`/api/v1/config`, {
            responseType: 'arraybuffer'
        })
            .then(res => {
                const resp = res.data;
                this.setState({ config: domopool_pb.Config.deserializeBinary(resp).toObject() });
                console.debug({ configViewOut: this.state.config });
                this.setState({ ipDisabled: this.state.config.network.dhcp });
            });
    }

    render() {
        if (this.state.config !== undefined) {
            if (this.state.config.global.displayStartup === undefined) {
                this.setState({ config: { global: { displayStartup: false } } });
            }

            return (
                <Card>
                    <CardContent>
                        <Formik
                            enableReinitialize
                            initialValues={{
                                global: this.state.config.global,
                            }
                            }
                            onSubmit={(values, actions) => {
                                // dataprovider.post(this.state.config.toArrayBuffer(),
                                //     {
                                //         responseType: 'arraybuffer',
                                //         headers: { 'Content-Type': 'application/octet-stream' }
                                //     }
                                // ).then(function (response) {
                                //     console.log(response)
                                // }).catch(function (response) {
                                //     console.log(response)
                                // })
                                alert(JSON.stringify(values, null, 2))
                            }}
                        >
                            {props => (

                                <Form>
                                    <GlobalForm {...props} />
                                    <br />
                                    <Button variant="contained" color="primary" type="submit">
                                        Submit
                                    </Button>
                                </Form>

                            )}
                        </Formik >
                        <Formik
                            enableReinitialize
                            initialValues={{
                                network: this.state.config.network,
                            }
                            }
                            onSubmit={(values, actions) => {
                                dataprovider.post(this.state.config.toArrayBuffer(),
                                    {
                                        responseType: 'arraybuffer',
                                        headers: { 'Content-Type': 'application/octet-stream' }
                                    }
                                ).then(function (response) {
                                    console.log(response)
                                })
                                    .catch(function (response) {
                                        console.log(response)
                                    })
                            }}
                        >
                            {props => (

                                <Form>
                                    <NetworkForm {...props} />
                                    <br />
                                    <Button variant="contained" color="primary" type="submit">
                                        Submit
                                    </Button>
                                </Form>

                            )}
                        </Formik >
                        <TempForm config={this.state.config.sensors} />
                    </CardContent>
                </Card>
            )
        } else {
            return (
                <Card>
                    <CardContent>
                        <LinearProgress />
                    </CardContent>
                </Card>
            )
        }
    };
};
