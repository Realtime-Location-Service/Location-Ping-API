package org.rls.rest.controller.location;

import org.rls.rest.model.Location;
import org.rls.rest.service.LocationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
public class LocationController {

    @Autowired
    private LocationService locationService;

    @PostMapping(path = "location/{object-id}")
    public ResponseEntity<Location> saveLocation(@PathVariable("object-id") Integer objectId, @RequestBody LocationApiInput location) {
        Location locationToSave = new Location();
        locationToSave.setObjectId(objectId);
        locationToSave.setLatitude(location.getLatitude());
        locationToSave.setLongitude(location.getLongitude());

        return new ResponseEntity<>(locationService.save(locationToSave), HttpStatus.CREATED);
    }

    @GetMapping(path = "location/{object-id}")
    public ResponseEntity<Location> getLocation(@PathVariable("object-id") Integer objectId) {
        Location location = locationService.findByObjectId(objectId);
        if (location != null) {
            return new ResponseEntity<>(location, HttpStatus.OK);
        }
        return new ResponseEntity<>(HttpStatus.NOT_FOUND);
    }
}
