import React from 'react';
import Relay from 'react-relay';
import PersonBalance from './PersonBalance';

class App extends React.Component {
  render() {
    return (
        <div>
            <section className="bounded">
                <div className="three-quarter-container bounded no-border">
                    <div className="span4">
                        <h3 className="ribbon-header">I currently owe</h3>
                        <div className="balances">
                            {
                                this.props.me.roommates.map(function(roommate) {
                                    return <PersonBalance key={roommate.id} person={roommate} />
                                })
                            }
                        </div>
                    </div>
                    <div className="span4 offset1">
                        <h3 className="ribbon-header">My roommates owe me</h3>
                        <div className="balances">
                            {
                                this.props.me.roommates.map(function(roommate) {
                                    return <PersonBalance key={roommate.id} person={roommate} />
                                })
                            }
                        </div>
                    </div>
                    <div className="clearfix"></div>
                </div>
            </section>

            <section className="bounded">
                <div className="three-quarter-container bounded no-border dark">
                    <div className='add-new-link'>
                        <a href="/expenses/new">Add Expense</a>
                        <a href="/expenses">All Expenses</a>
                    </div>
                    <h3 className='ribbon-header'>Recent Expenses</h3>
                    <div className='expenses'>
                        <div className='expense'>
                            <div className='expense-date'>
                                Nov 09
                            </div>
                            <div className='expense-amount'>
                                $1.00
                                <div className='expense-share'>
                                    your share is
                                    <span className='currency'>$1.00</span>
                                </div>
                            </div>
                            <div className='expense-description'>
                                Allison
                                <span className='expense-small'>paid for</span>
                                This is a test
                                <span className='expense-notes'>
                                    Seriously Just a Test
                                </span>
                                <span className='expense-small'>
                                    <a href="/expenses/1438" data-confirm="Are you sure you want to permanently delete this expense?" data-method="delete" rel="nofollow">delete</a>
                                </span>
                            </div>
                            <div className='clearfix'></div>
                        </div>
                        <div className='clearfix'></div>
                    </div>
                </div>
            </section>
        </div>
    );
  }
}

export default Relay.createContainer(App, {
  fragments: {
    me: () => Relay.QL`
        fragment on Person {
            name
            roommates {
                id
                ${PersonBalance.getFragment('person')}
            }
        }
    `,
  },
});
