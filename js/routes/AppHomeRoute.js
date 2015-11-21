import Relay from 'react-relay';

export default class extends Relay.Route {
  static queries = {
    latestPost: () => Relay.QL`
        query {
            latestPost
        }
    `,
  };
  static routeName = 'AppHomeRoute';
}
