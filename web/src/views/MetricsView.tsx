import * as React from "react";
import dataprovider from '../dataprovider/dataprovider';
import {
    LinearProgress,
    Card,
    CardContent,
} from "@material-ui/core";

var domopool_pb = require('../proto/domopool_pb');

export default class MetricsView extends React.Component {
    constructor(
        private config: any
    ) { super(config); };

    state = {
        config: this.config,
    }

    componentDidMount() {
        dataprovider.get(`/api/v1/config`, {
            responseType: 'arraybuffer'
        })
            .then(res => {
                const resp = res.data;
                this.setState({ config: domopool_pb.Config.deserializeBinary(resp).toObject() });
                console.log(this.state.config);
            })
    }

    render() {
        if (this.state.config.metrics !== undefined) {
            return (
                <Card>
                    <CardContent>
                        {this.state.config.metrics.twater}
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
