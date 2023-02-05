package com.unigrad.funiverseworkspacemanagementservice;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.net.InetAddress;
import java.net.UnknownHostException;

@RestController
@RequestMapping("/{tenant-id}/workspace-management")
public class WorkspaceManagementController {
    @GetMapping
    public ResponseEntity<String> getStringHello() throws UnknownHostException {
        return ResponseEntity.ok("Workspace Management: "+ InetAddress.getLocalHost().getHostName());
    }
}
