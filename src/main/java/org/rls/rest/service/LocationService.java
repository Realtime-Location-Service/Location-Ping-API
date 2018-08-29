package org.rls.rest.service;

import org.rls.rest.model.Location;

public interface LocationService {
    Location save(Location location);

    Location findByObjectId(Integer objectId);
}
