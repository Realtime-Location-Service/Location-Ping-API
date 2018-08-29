package org.rls.rest.controller;

import org.rls.rest.model.Location;
import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

@RunWith(SpringJUnit4ClassRunner.class)
@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
public class LocationControllerTest {

    @Autowired
    private TestRestTemplate template;

    @Before
    public void setup() {
    }

    @Test
    public void testLocationShouldBeCreated_WithMandatoryFields() {
        int objectId = 101;
        ResponseEntity<Location> locationFoundInCreate = createLocationThroughApi(getHttpEntity("{" +
                "\"latitude\": -21.345678" +
                ",\"longitude\": 12.98765" +
                "}"), objectId);

        ResponseEntity<Location> locationFoundInGet = template.getForEntity("/location/" + objectId, Location.class);
        Assert.assertEquals(locationFoundInCreate.getBody().getId(), locationFoundInGet.getBody().getId());
        Assert.assertEquals(-21.345678, locationFoundInGet.getBody().getLatitude(), 1e-6);
        Assert.assertEquals(12.98765, locationFoundInGet.getBody().getLongitude(), 1e-6);
    }

    @Test
    public void testLocationShouldNotBeCreated_WithoutMandatoryFields() {
        int objectId = 201;
        HttpEntity<Object> location = getHttpEntity("{" +
                "\"latitude\": -21.345678" +
                "}");
        ResponseEntity<Location> locationFoundInCreate = template.postForEntity("/location/" + objectId, location, Location.class);
        Assert.assertEquals(400, locationFoundInCreate.getStatusCode().value());
        Assert.assertNull(locationFoundInCreate.getBody().getId());
    }

    @Test
    public void testLocationShouldBeUpdated() {
        int objectId = 101;
        ResponseEntity<Location> locationFoundInCreate = createLocationThroughApi(getHttpEntity("{" +
                "\"latitude\": -21.345678" +
                ",\"longitude\": 12.98765" +
                "}"), objectId);

        ResponseEntity<Location> locationFoundInGet = template.getForEntity("/location/" + objectId, Location.class);
        Assert.assertEquals(locationFoundInCreate.getBody().getId(), locationFoundInGet.getBody().getId());
        Assert.assertEquals(-21.345678, locationFoundInGet.getBody().getLatitude(), 1e-6);
        Assert.assertEquals(12.98765, locationFoundInGet.getBody().getLongitude(), 1e-6);

        HttpEntity<Object> locationToUpdate = getHttpEntity("{" +
                "\"id\": 10012," + /* Even though we send id, it should be ignored */
                ",\"latitude\": 34.908765" +
                ",\"longitude\": 54.128712" +
                "}");
        template.put("/location/" + objectId, locationToUpdate);

        locationFoundInGet = template.getForEntity("/location/" + objectId, Location.class);
        Assert.assertEquals(locationFoundInCreate.getBody().getId(), locationFoundInGet.getBody().getId());
        Assert.assertEquals(34.908765, locationFoundInGet.getBody().getLatitude(), 1e-6);
        Assert.assertEquals(54.128712, locationFoundInGet.getBody().getLongitude(), 1e-6);
    }

    private ResponseEntity<Location> createLocationThroughApi(HttpEntity<Object> location, Integer objectId) {
        ResponseEntity<Location> locationFoundInCreate = template.postForEntity("/location/" + objectId, location, Location.class);
        Assert.assertEquals(201, locationFoundInCreate.getStatusCode().value());
        Assert.assertNotNull(locationFoundInCreate.getBody().getId());
        return locationFoundInCreate;
    }

    private HttpEntity<Object> getHttpEntity(Object body) {
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        return new HttpEntity<Object>(body, headers);
    }
}
