import React from 'react';
import Relay from 'react-relay';

class App extends React.Component {
  render() {
    return (
        <section className="bounded">
            <div className="three-quarter-container bounded no-border">
                <div className="span4">
                    <h3 className="ribbon-header">I currently owe</h3>
                    <div className="balances"></div>
                </div>
                <div className="span4 offset1">
                    <h3 className="ribbon-header">My roommates owe me</h3>
                    <div className="balances">
                        <div className='balance'>
                            <div className='span1 offset1'>
                                <span className='balance-name'> Allison McFarland </span>
                            </div>
                            <div className='span1'>
                                <span className='balance-amount'> $928.00 </span>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="clearfix"></div>
            </div>
        </section>
    );
  }
}

export default Relay.createContainer(App, {
  fragments: {
    me: () => Relay.QL`
        fragment on Person {
            name
            roommates {
                name
            }
        }
    `,
  },
});
