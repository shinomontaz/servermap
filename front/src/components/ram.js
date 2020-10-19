import React, { Component } from "react";
import { Icon } from 'semantic-ui-react'

class Ram extends Component {
  constructor(props) {
    super(props);
  }

  render () {
    var {maxRam, current, total} = this.props
    var chips = [];

    if (!maxRam) {
      maxRam = total
    }

    const maxItems = Math.ceil( ( total / maxRam ) * 5 ) ;
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
        return <Icon key={index} name='microchip' style={{ color: value }}/>
      })
    );
  }
}

export default Ram;
