import React, { useEffect } from 'react';

import { NavbarUserSettings } from '../components/Navbar';
import { Outlet } from 'react-router';
import styled from 'styled-components';

type Props = {};

const RouteContainer = styled.div`
  display: flex;
  height: 100%;
`;

const RouterOutlet = styled.main`
  padding: 1em;
  width: 100%;
  height: 100%;
  overflow-y: auto;
`;

const UserSettingsRoute: React.FC<Props> = ({}) => {
  useEffect(() => {}, []);

  return (
    <RouteContainer>
      <NavbarUserSettings />
      <RouterOutlet>
        <Outlet />
      </RouterOutlet>
    </RouteContainer>
  );
};

export default UserSettingsRoute;
