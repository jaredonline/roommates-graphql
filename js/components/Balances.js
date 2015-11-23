import React from 'react';
import Relay from 'react-relay';
import PersonBalance from './Balances/Person';

class App extends React.Component {
  render() {
    return (
        <section className="bounded">
            <div className="three-quarter-container bounded no-border">
                <div className="span4">
                    <h3 className="ribbon-header">I currently owe</h3>
                    <div className="balances">
                        {
                            this.props.person.roommates.map(function(roommate) {
                                return <PersonBalance key={roommate.id} person={roommate} />
                            })
                        }
                    </div>
                </div>
                <div className="span4 offset1">
                    <h3 className="ribbon-header">My roommates owe me</h3>
                    <div className="balances">
                        {
                            this.props.person.roommates.map(function(roommate) {
                                return <PersonBalance key={roommate.id} person={roommate} />
                            })
                        }
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
    person: () => Relay.QL`
        fragment on Person {
            roommates {
                id
                ${PersonBalance.getFragment('person')}
            }
        }
    `,
  },
});
