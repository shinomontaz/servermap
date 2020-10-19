import React, { Component } from "react";
import { Icon } from 'semantic-ui-react'

class Cpu extends Component {
  constructor(props) {
    super(props);
    this.state = {
      open: false,
    };
  }

  render () {
    var {maxCpu, current, total} = this.props
    var chips = [];

    if (!maxCpu) {
      maxCpu = total
    }

    const maxItems = Math.ceil( ( total / maxCpu ) * 5 ) ;

    const busyItems = Math.floor( ( current / total ) * maxItems ) ;
    const topBusyItems = Math.ceil( ( current / total ) * maxItems ) ;

    for (var i=0; i<maxItems; i++) {
      chips[i] = "green";
      if (current > total ) {
        chips[i] = "red";
      }
      if (i < topBusyItems && i >= busyItems) {
        chips[i] = "olive";
      }
      if (i >= topBusyItems) {
        chips[i] = "grey";
      }
    }

    return(
      chips.map((value, index) => {
        return <Icon key={index} name='th' style={{ color: value }}/>
      })
    );
  }
}

export default Cpu;
