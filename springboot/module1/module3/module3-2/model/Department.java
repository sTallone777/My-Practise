package model;

/**
 * Table: department
 */
public class Department {
    /**
     * Column: id
     */
    private Integer id;

    /**
     * Column: name
     * Remark: 部门名称
     */
    private String name;

    /**
     * Column: parentId
     */
    private Integer parentid;

    /**
     * Column: depPath
     */
    private String deppath;

    /**
     * Column: enabled
     */
    private Boolean enabled;

    /**
     * Column: isParent
     */
    private Boolean isparent;

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name == null ? null : name.trim();
    }

    public Integer getParentid() {
        return parentid;
    }

    public void setParentid(Integer parentid) {
        this.parentid = parentid;
    }

    public String getDeppath() {
        return deppath;
    }

    public void setDeppath(String deppath) {
        this.deppath = deppath == null ? null : deppath.trim();
    }

    public Boolean getEnabled() {
        return enabled;
    }

    public void setEnabled(Boolean enabled) {
        this.enabled = enabled;
    }

    public Boolean getIsparent() {
        return isparent;
    }

    public void setIsparent(Boolean isparent) {
        this.isparent = isparent;
    }
}