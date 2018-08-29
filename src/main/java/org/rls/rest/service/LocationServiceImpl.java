package org.rls.rest.service;

import org.rls.rest.model.Location;
import org.rls.rest.repository.LocationRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class LocationServiceImpl implements LocationService {

    @Autowired
    LocationRepository locationRepository;

    public Location save(Location location) {
        return locationRepository.save(location);
    }

    public Location findByObjectId(Integer objectId) {
        return locationRepository.findByObjectId(objectId);
    }
}