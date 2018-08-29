package org.rls.rest.controller.location;

import javax.validation.constraints.NotNull;
import java.io.Serializable;

public class LocationApiInput implements Serializable {

    @NotNull
    Double latitude;

    @NotNull
    Double longitude;

    public Double getLatitude() {
        return latitude;
    }

    public void setLatitude(Double latitude) {
        this.latitude = latitude;
    }

    public Double getLongitude() {
        return longitude;
    }

    public void setLongitude(Double longitude) {
        this.longitude = longitude;
    }
}
