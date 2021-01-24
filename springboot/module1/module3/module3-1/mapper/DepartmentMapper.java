package mapper;

import java.util.List;
import model.Department;
import model.DepartmentExample;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.session.RowBounds;
import org.springframework.stereotype.Repository;

@Repository
public interface DepartmentMapper {
    /**
     * @mbg.generated generated automatically, do not modify!
     */
    long countByExample(DepartmentExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int deleteByExample(DepartmentExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int deleteByPrimaryKey(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int insert(Department record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int insertSelective(Department record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    List<Department> selectByExample(DepartmentExample example, RowBounds rowBounds);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    List<Department> selectByExample(DepartmentExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    Department selectByPrimaryKeyWithLock(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    Department selectByPrimaryKey(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByExampleSelective(@Param("record") Department record, @Param("example") DepartmentExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByExample(@Param("record") Department record, @Param("example") DepartmentExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByPrimaryKeySelective(Department record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByPrimaryKey(Department record);
}